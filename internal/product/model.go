package product

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"unique:not null" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Stock       int       `json:"stock" validate:"required,gt=0"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_aat"`
	UpdatedAt   time.Time `json:"updated_at"`
}
