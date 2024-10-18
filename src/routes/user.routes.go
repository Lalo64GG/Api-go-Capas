package routes

import (
	"github.com/gin-gonic/gin"

	"firstAPI-go/src/controllers"
	"firstAPI-go/src/services"
)

func UserRoutes(router *gin.RouterGroup, userService *services.UserService) {

	userController := controllers.NewUserController(userService)

	router.POST("/", userController.CreateUser)
}