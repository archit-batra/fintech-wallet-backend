package user

import "strconv"

type Service struct {
	users map[string]User
	idSeq int
}

func NewService() *Service {
	return &Service{
		users: make(map[string]User),
	}
}

func (s *Service) CreateUser(name, email string) User {
	s.idSeq++

	id := strconv.Itoa(s.idSeq)

	u := User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	s.users[id] = u
	return u
}

func (s *Service) GetUser(id string) (User, bool) {
	u, ok := s.users[id]
	return u, ok
}
