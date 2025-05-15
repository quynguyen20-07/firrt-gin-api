package service

import (
	"github.com/gin/api/dto"
	"github.com/gin/api/models"
	"github.com/gin/api/repository"
)

type ProductService interface {
	Create(dto.CreateUserDTO) error
	GetAll() ([]models.User, error)
}

type productService struct {
	repo repository.UserRepository
}

func NewProductService(repo repository.UserRepository) ProductService {
	return &productService{repo}
}

func (s *productService) Create(data dto.CreateUserDTO) error {
	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	return s.repo.Create(&user)
}

func (s *productService) GetAll() ([]models.User, error) {
	return s.repo.FindAll()
}
