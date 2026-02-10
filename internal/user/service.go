package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(name, email string) (User, error) {
	return s.repo.CreateUser(name, email)
}

func (s *Service) GetUser(id string) (User, error) {
	return s.repo.GetUser(id)
}
