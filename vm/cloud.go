package services

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

import (
	"errors"
	"sync"
)

type InMemoryUserService struct {
	users map[string]*models.User
	mu    sync.Mutex
}

func NewInMemoryUserService() *InMemoryUserService {
	return &InMemoryUserService{
		users: make(map[string]*models.User),
	}
}

func (s *InMemoryUserService) CreateUser(user *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	s.users[user.ID] = user
	return nil
}

