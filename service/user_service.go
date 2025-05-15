package service

import (
	"errors"
	"strings"

	"github.com/gin/api/dto"
	"github.com/gin/api/models"
	"github.com/gin/api/repository"
)

type UserService interface {
	Create(dto.CreateUserDTO) error
	GetAll() ([]models.User, error)
	GetUserByEmail(email string) ([]models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Create(data dto.CreateUserDTO) error {
	// 1. Kiểm tra trùng email (giả sử repo có hàm FindByEmail)
	existing, _ := s.repo.FindOneByEmail(strings.ToLower(data.Email))
	if existing != nil {
		return errors.New("email already exists")
	}

	// 2. Hash mật khẩu (giả sử tạm thời không hash)
	user := models.User{
		Name:     data.Name,
		Email:    strings.ToLower(data.Email),
		Password: data.Password, // TODO: hash sau
	}

	return s.repo.Create(&user)
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByEmail(email string) ([]models.User, error) {
	user, err := s.repo.FindOneByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return []models.User{*user}, nil
}
