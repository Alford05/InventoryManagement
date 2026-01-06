package order

import "time"

type Order struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	CustomerID uint        `json:"customer_id" gorm:"not null"`
	Status     Status      `json:"status" gorm:"type:varchar(20);not null"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	CreatedAt  time.Time   `json:"created_at"`
}

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" validate:"required,gt=0"`
	UnitPrice float64 `json:"unit_price"`
}
