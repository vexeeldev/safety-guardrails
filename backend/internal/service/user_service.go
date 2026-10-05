package service

import (
	"github.com/google/uuid"
	"github.com/luki/safety-guardrails-backend/internal/repository"
	"github.com/luki/safety-guardrails-backend/models"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return s.repo.FindByID(id)
}
