package repository

import (
	"WebParkir/apps/api/internal/domain"

	"gorm.io/gorm"
)

type VehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) *VehicleRepository {
	return &VehicleRepository{db}
}

func (r *VehicleRepository) Create(vehicle domain.Vehicle) (*domain.Vehicle, error) {
	err := r.db.Create(&vehicle).Error

	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (r *VehicleRepository) GetAll() ([]domain.Vehicle, error) {
	var vehicles []domain.Vehicle

	err := r.db.Find(&vehicles).Error

	return vehicles, err
}

func (r *VehicleRepository) FindByID(id uint) (*domain.Vehicle, error) {
	var vehicle domain.Vehicle

	err := r.db.First(&vehicle, id).Error

	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (r *VehicleRepository) FindByLicensePlate(licensePlate string) ([]domain.Vehicle, error) {
	var vehicles []domain.Vehicle

	err := r.db.
		Where("license_plate ILIKE ?", "%"+licensePlate+"%").
		Find(&vehicles).Error

	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (r *VehicleRepository) Update(vehicle domain.Vehicle) (*domain.Vehicle, error) {
	err := r.db.Save(&vehicle).Error

	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (r *VehicleRepository) Delete(id uint) error {
	err := r.db.Delete(&domain.Vehicle{}, id).Error

	return err
}
