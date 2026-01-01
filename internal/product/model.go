package product

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"unique:not null" validate:"required"`
	Description string  `validate:"required"`
	Price       float64 `validate:"required,gt=0"`
	Stock       int     `validate:"required,gt=0"`
	CategoryID  uint    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
