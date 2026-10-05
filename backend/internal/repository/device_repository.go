package repository

import (
	"github.com/google/uuid"
	"github.com/luki/safety-guardrails-backend/models"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{db: db}
}

func (r *DeviceRepository) Create(device *models.Device) error {
	return r.db.Create(device).Error
}

func (r *DeviceRepository) FindByID(id uuid.UUID) (*models.Device, error) {
	var device models.Device
	err := r.db.First(&device, "id = ?", id).Error
	return &device, err
}

func (r *DeviceRepository) FindByUserID(userID uuid.UUID) ([]models.Device, error) {
	var devices []models.Device
	err := r.db.Where("user_id = ?", userID).Find(&devices).Error
	return devices, err
}
