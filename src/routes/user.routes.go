package routes

import (
	"github.com/gin-gonic/gin"

	"firstAPI-go/src/controllers" 
	"firstAPI-go/src/services"    
)

//? UserRoutes define las rutas relacionadas con los usuarios
//* Este método recibe un grupo de rutas (RouterGroup) y un UserService para manejar la lógica de negocio
func UserRoutes(router *gin.RouterGroup, userService *services.UserService) {

	//* Se crea un nuevo controlador de usuarios pasando el servicio de usuario
	userController := controllers.NewUserController(userService)

	//* Define la ruta POST para crear un usuario
	//* Cuando se haga una solicitud POST a esta ruta, se llamará al método CreateUser del controlador
	router.POST("/", userController.CreateUser)
}
