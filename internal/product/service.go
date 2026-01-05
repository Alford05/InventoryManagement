package product

import "InventoryManagement/pkg/validation"

type Service interface {
	Create(product *Product) error
	List(filters map[string]interface{}) ([]Product, error)
	Get(id uint) (*Product, error)
	Update(id uint, product *Product) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(p *Product) error {
	if err := validation.Validate.Struct(p); err != nil {
		return err
	}
	return s.repo.Create(p)
}

func (s *service) List(filters map[string]interface{}) ([]Product, error) {
	return s.repo.FindAll(filters)
}

func (s *service) Get(id uint) (*Product, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(id uint, p *Product) error {
	return s.repo.Update(id, p)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
