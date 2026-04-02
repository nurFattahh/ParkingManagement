package services

import (
	"WebParkir/apps/api/internal/domain"
	"WebParkir/apps/api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) Register(req domain.RegisterRequest) (*domain.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		return nil, err
	}

	user := domain.User{
		Username: req.Username,
		Password: string(hash),
		FullName: req.FullName,
		Phone:    req.Phone,
		Email:    req.Email,
		Role:     "admin",
	}

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(req domain.LoginRequest) (*domain.User, error) {

	user, err := s.userRepo.FindByUsername(req.Username)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) GetUsers() ([]domain.User, error) {
	return s.userRepo.GetAll()
}

func (s *AuthService) AddUser(req domain.RegisterRequest) (*domain.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		return nil, err
	}

	user := domain.User{
		Username: req.Username,
		Password: string(hash),
		Role:     "user",
	}

	return s.userRepo.AddUser(user)
}
