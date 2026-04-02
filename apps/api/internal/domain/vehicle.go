package domain

import "time"

type Vehicle struct {
	ID           uint           `gorm:"primaryKey"`
	VehicleType  string         `gorm:"not null"`
	LicensePlate string         `gorm:"unique;not null"`
	EntryTime    time.Time      `gorm:"not null"`
	ExitTime     *time.Time     `gorm:"default:null"`
	Duration     *time.Duration `gorm:"default:null"`
	Tariff       *float64       `gorm:"default:null"`
	Status       string         `gorm:"not null"`
}

type VehicleEntryRequest struct {
	LicensePlate string `json:"license_plate" validate:"required"`
}

type VehicleExitRequest struct {
	LicensePlate string `json:"license_plate" validate:"required"`
}

type AddVehicleRequest struct {
	LicensePlate string `json:"license_plate" validate:"required"`
	VehicleType  string `json:"vehicle_type" validate:"required"`
}
