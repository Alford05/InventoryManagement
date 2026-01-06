package product

import "gorm.io/gorm"

type Repository interface {
	Create(product *Product) error
	FindAll(query *ListProductsQuery) ([]Product, int64, error)
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

func (r *repository) FindAll(query *ListProductsQuery) ([]Product, int64, error) {
	var products []Product
	dbQuery := r.db.Model(&Product{})

	if query.Category != "" {
		dbQuery = dbQuery.Where("category_id = ?", query.Category)
	}
	if query.MinPrice > 0 {
		dbQuery = dbQuery.Where("price >= ?", query.MinPrice)
	}
	if query.MaxPrice > 0 {
		dbQuery.Where("price <= ?", query.MaxPrice)
	}
	if query.Search != "" {
		search := "%" + query.Search + "%"
		dbQuery = dbQuery.Where("name ILIKE ? OR description ILIKE ?", search, search)
	}

	var total int64
	dbQuery.Count(&total)

	offset := (query.Page - 1) * query.PageSize
	dbQuery = dbQuery.Offset(offset).Limit(query.PageSize)

	err := dbQuery.Find(&products).Error
	return products, total, err
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
