package category

import "gorm.io/gorm"

type Repository interface {
	Create(category *Category) error
	FindAll() ([]Category, error)
	FindByID(id uint) (*Category, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Category{})
	return &repository{db}
}

func (r *repository) Create(c *Category) error {
	return r.db.Create(c).Error
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *repository) FindByID(id uint) (*Category, error) {
	var category Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Category{}, id).Error
}
