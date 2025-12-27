package category

type Service interface {
	Create(category *Category) error
	List() ([]Category, error)
	Get(id uint) (*Category, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(c *Category) error {
	return s.repo.Create(c)
}

func (s *service) List() ([]Category, error) {
	return s.repo.FindAll()
}

func (s *service) Get(id uint) (*Category, error) {
	return s.repo.FindByID(id)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
