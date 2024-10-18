package services

import (
	"database/sql"  //* Paquete estándar para interactuar con bases de datos SQL
	"log"          
	"firstAPI-go/src/models"  
)

//? UserService es una estructura que contiene la conexión a la base de datos.
//* Esta estructura será utilizada para interactuar con la tabla de usuarios y manejar
//* las operaciones relacionadas con los usuarios, como la creación, actualización, eliminación, etc.
type UserService struct {
	DB *sql.DB  //* DB es una instancia de *sql.DB, que representa una conexión activa a la base de datos.
}

//? NewUserService es un constructor que inicializa una nueva instancia de UserService.
//* Toma como parámetro una conexión activa a la base de datos (*sql.DB) y retorna una instancia de UserService.
//* El propósito de esta función es inyectar la dependencia de la base de datos en el servicio de usuarios.
func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}  //* Retorna una nueva instancia de UserService con la conexión a la base de datos
}

//? CreateUser es un método del servicio de usuarios que se encarga de insertar un nuevo registro de usuario
//* en la base de datos. Recibe un objeto `models.User` que contiene la información del usuario a insertar.
//* Si ocurre algún error durante la inserción, lo registra en el log y lo retorna para que pueda ser manejado
//* por capas superiores (como controladores). Si la inserción es exitosa, registra un mensaje de éxito.
func (s *UserService) CreateUser(user models.User) error {
	//* Preparamos la consulta SQL para insertar un nuevo usuario en la tabla `users`.
	//* Utilizamos placeholders `?` para evitar inyecciones SQL al permitir el paso de los valores de manera segura.
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"

	//* Ejecutamos la consulta SQL usando `Exec`, que ejecuta una sentencia y devuelve el resultado.
	//* Aquí pasamos los valores correspondientes (nombre, email y contraseña) usando `user.Name`, `user.Email` y `user.Password`.
	_, err := s.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		//* Si ocurre un error durante la ejecución de la consulta, lo registramos en el log
		//* y devolvemos el error para que pueda ser manejado externamente.
		log.Printf("Error al crear el usuario: %v", err)
		return err
	}

	//* Si la consulta fue exitosa, registramos un mensaje en el log indicando que el usuario fue creado correctamente.
	log.Println("Usuario creado correctamente")
	return nil
}
