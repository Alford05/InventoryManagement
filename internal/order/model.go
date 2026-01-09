package order

import "time"

type OrderStatus string

const (
	StatusPending  OrderStatus = "pending"
	StatusPaid     OrderStatus = "paid"
	StatusShipped  OrderStatus = "shipped"
	StatusCanceled OrderStatus = "canceled"
)

type Order struct {
	ID         uint        `gorm:"primaryKey"`
	CustomerID uint        `gorm:"not null;index"`
	Status     OrderStatus `gorm:"type:varchar(20);default:'pending'"`
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	UnitPrice float64 `gorm:"not null"`
}

type UpdateOrderStatusRequest struct {
	Status OrderStatus `json:"status" validate:"required,oneof=pending paid shipped canceled"`
}
