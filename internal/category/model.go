package category

import "time"

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique:not null" validate:"required"`
	CreatedAt time.Time
}
