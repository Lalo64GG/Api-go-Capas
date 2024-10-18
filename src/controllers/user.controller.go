package controllers

import (
	"firstAPI-go/src/models"    
	"firstAPI-go/src/services"  // Importamos el servicio UserService, que contiene la lógica para interactuar con la base de datos
	"github.com/gin-gonic/gin" 
	"net/http"                 
)

//? UserController es una estructura que contiene una referencia a UserService.
//* Este controlador es responsable de manejar las solicitudes HTTP relacionadas con los usuarios.
type UserController struct {
	UserService *services.UserService  //* UserService es una referencia al servicio de usuarios, que encapsula la lógica de negocio para manejar usuarios.
}

//? NewUserController es un constructor que crea una nueva instancia de UserController.
//* Este constructor recibe una instancia de UserService que será utilizada por el controlador para realizar operaciones relacionadas con usuarios.
//* El patrón de inyección de dependencias asegura que el controlador y el servicio estén desacoplados, facilitando las pruebas y el mantenimiento.
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}  //* Devuelve una nueva instancia de UserController con el servicio de usuarios inyectado
}

//? CreateUser es un método del controlador que maneja las solicitudes HTTP POST para crear un nuevo usuario.
//* Este método recibe el contexto de la solicitud HTTP (`gin.Context`) y utiliza el servicio de usuarios (`UserService`) para insertar un nuevo usuario en la base de datos.
func (ctrl *UserController) CreateUser(ctx *gin.Context) {
	var user models.User  //* Creamos una instancia del modelo User para almacenar los datos del cuerpo de la solicitud

	//* Vinculamos los datos JSON enviados en la solicitud HTTP al modelo User.
	//* `ShouldBindJSON` vincula el JSON al struct Go (en este caso, `user`) y devuelve un error si los datos no son válidos o están incompletos.
	if err := ctx.ShouldBindJSON(&user); err != nil {
		//* Si ocurre un error durante el enlace (por ejemplo, si falta algún campo obligatorio), respondemos con un código de estado 400 (Bad Request)
		//* y un mensaje que indica que los datos son inválidos.
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	//* Si los datos son válidos, llamamos al servicio `UserService` para crear el usuario en la base de datos.
	err := ctrl.UserService.CreateUser(user)
	if err != nil {
		//* Si ocurre un error en el servicio al intentar crear el usuario (por ejemplo, un error en la base de datos),
		//* respondemos con un código de estado 500 (Internal Server Error) y un mensaje que indica el fallo.
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}

	//* Si el usuario fue creado correctamente, respondemos con un código de estado 200 (OK)
	//* y un mensaje que indica que la operación fue exitosa.
	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario creado correctamente"})
}
