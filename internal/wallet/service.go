package wallet

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateWallet(userID int) error {
	return s.repo.CreateWallet(userID)
}

func (s *Service) AddMoney(userID int, amount int64) error {
	return s.repo.AddBalance(userID, amount)
}

func (s *Service) GetWallet(userID int) (Wallet, error) {
	return s.repo.GetWallet(userID)
}
