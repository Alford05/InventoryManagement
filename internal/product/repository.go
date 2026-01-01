package product

import "gorm.io/gorm"

type Repository interface {
	Create(product *Product) error
	FindAll(filters map[string]interface{}) ([]Product, error)
	FindByID(id uint) (*Product, error)
	Update(id uint, product *Product) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Product{})
	return &repository{db}
}

func (r *repository) Create(p *Product) error {
	return r.db.Create(p).Error
}

func (r *repository) FindAll(filters map[string]interface{}) ([]Product, error) {
	var products []Product
	err := r.db.Where(filters).Find(&products).Error
	return products, err
}

func (r *repository) FindByID(id uint) (*Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *repository) Update(id uint, p *Product) error {
	return r.db.Model(&Product{}).Where("id = ?", id).Updates(p).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Product{}, id).Error
}
