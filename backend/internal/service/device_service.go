package service

import (
	"github.com/google/uuid"
	"github.com/luki/safety-guardrails-backend/internal/repository"
	"github.com/luki/safety-guardrails-backend/models"
)

type DeviceService struct {
	repo *repository.DeviceRepository
}

func NewDeviceService(repo *repository.DeviceRepository) *DeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) CreateDevice(device *models.Device) error {
	return s.repo.Create(device)
}

func (s *DeviceService) GetDeviceByID(id uuid.UUID) (*models.Device, error) {
	return s.repo.FindByID(id)
}

func (s *DeviceService) GetDevicesByUserID(userID uuid.UUID) ([]models.Device, error) {
	return s.repo.FindByUserID(userID)
}
