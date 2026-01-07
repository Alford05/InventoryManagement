package order

import "gorm.io/gorm"

type Repository interface {
	WithTx(tx *gorm.DB) Repository
	Create(order *Order) error
	CreateItem(item *OrderItem) error
	FindProductForUpdate(productID uint) (*ProductSnapshot, error)
	UpdateProductStock(productID uint, newStock int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Order{}, &OrderItem{})
	return &repository{db: db}
}

func (r *repository) WithTx(tx *gorm.DB) Repository {
	return &repository{db: tx}
}

type ProductSnapshot struct {
	ID    uint
	Price float64
	Stock int
}

func (r *repository) Create(order *Order) error {
	return r.db.Create(order).Error
}

func (r *repository) CreateItem(item *OrderItem) error {
	return r.db.Create(item).Error
}

func (r *repository) FindProductForUpdate(productID uint) (*ProductSnapshot, error) {
	var p ProductSnapshot
	err := r.db.
		Table("products").
		Clauses(gorm.Expr("FOR UPDATE")).
		Where("id = ?", productID).
		First(&p).Error
	return &p, err
}

func (r *repository) UpdateProductStock(productID uint, newStock int) error {
	return r.db.
		Table("products").
		Where("id = ?", productID).
		Update("stock", newStock).Error
}
