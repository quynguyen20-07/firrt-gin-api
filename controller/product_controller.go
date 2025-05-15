package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin/api/dto"
	"github.com/gin/api/service"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(s service.ProductService) *ProductController {
	return &ProductController{s}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var input dto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.productService.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
