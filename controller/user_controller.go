package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin/api/dto"
	"github.com/gin/api/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{s}
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/users")
	group.POST("/", c.CreateUser)
	group.GET("/", c.GetAllUsers)
	group.GET("/:email", c.GetUserByEmail)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var input dto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.userService.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email") // 👈 Lấy `email` từ URL
	users, err := c.userService.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
