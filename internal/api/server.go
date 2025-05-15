package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin/api/controller"
	"github.com/gin/api/database"
	"github.com/gin/api/repository"
	"github.com/gin/api/service"
)

func StartServer() {

	database.Connect()
	r := gin.Default()

	userRepo := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(r)

	r.Run(":8080")
}
