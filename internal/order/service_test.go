package order_test

import (
	"testing"

	"InventoryManagement/internal/order"
	"InventoryManagement/internal/product"
	"InventoryManagement/internal/testutils"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrder_Success(t *testing.T) {
	db := testutils.SetupTestDB()

	// Seed product
	p := product.Product{
		Name:  "Phone",
		Price: 500,
		Stock: 10,
	}
	db.Create(&p)

	repo := order.NewRepository(db)
	service := order.NewService(db, repo)

	req := &order.CreateOrderRequest{
		CustomerID: 1,
		Items: []order.CreateOrderItemRequest{
			{ProductID: p.ID, Quantity: 2},
		},
	}

	o, err := service.CreateOrder(req)

	assert.NoError(t, err)
	assert.NotNil(t, o)
	assert.Equal(t, order.StatusPending, o.Status)

	// Verify stock deduction
	var updated product.Product
	db.First(&updated, p.ID)
	assert.Equal(t, 8, updated.Stock)
}

func TestCreateOrder_InsufficientStock(t *testing.T) {
	db := testutils.SetupTestDB()

	p := product.Product{
		Name:  "Laptop",
		Price: 1000,
		Stock: 1,
	}
	db.Create(&p)

	repo := order.NewRepository(db)
	service := order.NewService(db, repo)

	req := &order.CreateOrderRequest{
		CustomerID: 1,
		Items: []order.CreateOrderItemRequest{
			{ProductID: p.ID, Quantity: 2},
		},
	}

	o, err := service.CreateOrder(req)

	assert.Error(t, err)
	assert.Nil(t, o)

	// Verify stock unchanged
	var updated product.Product
	db.First(&updated, p.ID)
	assert.Equal(t, 1, updated.Stock)
}

func TestCancelOrder_RestoresStock(t *testing.T) {
	db := testutils.SetupTestDB()

	p := product.Product{
		Name:  "Tablet",
		Price: 300,
		Stock: 5,
	}
	db.Create(&p)

	repo := order.NewRepository(db)
	service := order.NewService(db, repo)

	// Create order
	req := &order.CreateOrderRequest{
		CustomerID: 1,
		Items: []order.CreateOrderItemRequest{
			{ProductID: p.ID, Quantity: 3},
		},
	}
	o, _ := service.CreateOrder(req)

	// Cancel order
	_, err := service.CancelOrder(o.ID)
	assert.NoError(t, err)

	var updated product.Product
	db.First(&updated, p.ID)
	assert.Equal(t, 5, updated.Stock)
}
