package order

import (
	"errors"

	"InventoryManagement/pkg/validation"

	"gorm.io/gorm"
)

type Service interface {
	CreateOrder(req *CreateOrderRequest) (*Order, error)
	GetOrder(id uint) (*Order, error)
	ListOrders(customerID uint) ([]Order, error)
}

type service struct {
	db   *gorm.DB
	repo Repository
}

func (s *service) GetOrder(id uint) (*Order, error) {
	var order Order
	err := s.db.Preload("Items").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *service) ListOrders(customerID uint) ([]Order, error) {
	var orders []Order
	query := s.db.Preload("Items")
	if customerID != 0 {
		query = query.Where("customer_id = ?", customerID)
	}
	err := query.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func NewService(db *gorm.DB, repo Repository) Service {
	return &service{db: db, repo: repo}
}

func (s *service) CreateOrder(req *CreateOrderRequest) (*Order, error) {
	// validate the request
	if err := validation.Validate.Struct(req); err != nil {
		return nil, err
	}

	var order Order

	err := s.db.Transaction(func(tx *gorm.DB) error {
		txRepo := s.repo.WithTx(tx)

		// Create Order
		order = Order{
			CustomerID: req.CustomerID,
			Status:     StatusPending,
		}
		if err := txRepo.Create(&order); err != nil {
			return err
		}

		// Process each order item
		for _, item := range req.Items {
			product, err := txRepo.FindProductForUpdate(item.ProductID)
			if err != nil {
				return err
			}
			if product.Stock < item.Quantity {
				return errors.New("insufficient stock")
			}

			// Deduct stock
			newStock := product.Stock - item.Quantity
			if err := txRepo.UpdateProductStock(product.ID, newStock); err != nil {
				return err
			}

			// Create order item
			orderItem := OrderItem{
				OrderID:   order.ID,
				ProductID: product.ID,
				Quantity:  item.Quantity,
				UnitPrice: product.Price,
			}

			if err := txRepo.CreateItem(&orderItem); err != nil {
				return err
			}
		}
		return nil // commit the transaction
	})
	if err != nil {
		return nil, err
	}

	return &order, nil
}
