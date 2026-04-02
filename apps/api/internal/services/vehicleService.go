package services

import (
	"WebParkir/apps/api/internal/domain"
	"WebParkir/apps/api/internal/repository"
	"errors"
	"time"
)

type VehicleService struct {
	vehicleRepo *repository.VehicleRepository
}

func NewVehicleService(vehicleRepo *repository.VehicleRepository) *VehicleService {
	return &VehicleService{vehicleRepo}
}

func (s *VehicleService) AddVehicle(vehicle domain.AddVehicleRequest) (*domain.Vehicle, error) {

	if vehicle.LicensePlate == "" {
		return nil, errors.New("license plate is required")
	}

	existing, err := s.vehicleRepo.FindByLicensePlate(vehicle.LicensePlate)
	if err == nil && existing != nil {
		return nil, errors.New("vehicle with this license plate already exists")
	}

	vehicleEntity := domain.Vehicle{
		LicensePlate: vehicle.LicensePlate,
		VehicleType:  vehicle.VehicleType,
		EntryTime:    time.Now(),
		Status:       "parked",
	}

	return s.vehicleRepo.Create(vehicleEntity)
}

func (s *VehicleService) GetAllVehicles() ([]domain.Vehicle, error) {
	vehicles, err := s.vehicleRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []domain.Vehicle

	for _, v := range vehicles {
		duration := time.Since(v.EntryTime)

		result = append(result, domain.Vehicle{
			ID:           v.ID,
			LicensePlate: v.LicensePlate,
			VehicleType:  v.VehicleType,
			EntryTime:    v.EntryTime,
			Status:       v.Status,
			Duration:     &duration, // atau format sendiri
		})
	}

	return result, nil
}

func (s *VehicleService) GetVehicleByID(id uint) (*domain.Vehicle, error) {
	return s.vehicleRepo.FindByID(id)
}

func (s *VehicleService) GetVehicleByLicensePlate(licensePlate string) (*domain.Vehicle, error) {
	vehicle, err := s.vehicleRepo.FindByLicensePlate(licensePlate)
	if err != nil {
		return nil, errors.New("vehicle not found")
	}

	// 🔥 hitung duration
	duration := time.Since(vehicle.EntryTime)
	vehicle.Duration = &duration

	return vehicle, nil
}

func (s *VehicleService) UpdateVehicle(vehicle domain.Vehicle) (*domain.Vehicle, error) {
	return s.vehicleRepo.Update(vehicle)
}

func (s *VehicleService) DeleteVehicle(id uint) error {
	return s.vehicleRepo.Delete(id)
}
