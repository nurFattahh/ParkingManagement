package repository

import (
	"WebParkir/apps/api/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user domain.User) (*domain.User, error) {

	err := r.db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*domain.User, error) {

	var user domain.User

	err := r.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetAll() ([]domain.User, error) {

	var users []domain.User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *UserRepository) AddUser(user domain.User) (*domain.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
