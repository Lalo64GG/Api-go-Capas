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

//? GetAllUsers es un método del servicio de usuarios que se encarga de obtener todos los usuarios
func (s *UserService) GetUsersAll() ([]models.User, error) {
    //* Corregimos la consulta SQL, ahora seleccionando los campos correctamente.
    query := "SELECT id, name, password FROM users"

    //* Ejecutamos la consulta
    rows, err := s.DB.Query(query)
    if err != nil {
        log.Printf("Error al traer los usuarios: %v", err)
        return nil, err
    }
    defer rows.Close() //* Cerramos las filas cuando ya no se necesiten.

    //* Creamos un slice para almacenar los usuarios
    var users []models.User

    //* Recorremos las filas obtenidas
    for rows.Next() {
        var user models.User
        //? Escaneamos cada fila en un objeto user
		//* La función rows.Next() recorre fila por fila los resultados obtenidos de la consulta SQL.
		//* En cada iteración, necesitamos extraer los datos de la fila actual y almacenarlos
		//* en una estructura adecuada para poder utilizarlos en nuestra aplicación.
		//* La función rows.Scan() nos permite asignar los valores de las columnas de la fila actual
		//* a las propiedades correspondientes de un objeto (en este caso, de tipo models.User).
		//* Específicamente, estamos mapeando las columnas "id", "name" y "password" a los campos
		//* user.ID, user.Name, y user.Password, respectivamente.
		//* Si el mapeo falla, se registra el error y se retorna para evitar procesar datos incompletos o incorrectos.
        if err := rows.Scan(&user.ID, &user.Name, &user.Password); err != nil {
            log.Printf("Error al escanear los usuarios: %v", err)
            return nil, err
        }
        users = append(users, user) //* Añadimos el usuario al slice
    }

    //* Verificamos si hubo errores en el recorrido de las filas
    if err := rows.Err(); err != nil {
        log.Printf("Error durante la iteración de filas: %v", err)
        return nil, err
    }

    log.Print("Se han obtenido los usuarios de manera exitosa")
    return users, nil //* Retornamos el slice de usuarios
}


//? GetUserByID es un método del servicio de usuarios que se encarga de obtener un usuario por su ID
func (s *UserService) GetUserByID(id int) (models.User, error) {
    //* Consulta SQL para obtener un usuario específico por su ID.
    query := "SELECT id, name, email, password FROM users WHERE id = ?"

    //* Ejecutamos la consulta para obtener una sola fila usando QueryRow(), ya que solo esperamos un resultado.
    row := s.DB.QueryRow(query, id)

    //* Creamos una variable para almacenar los datos del usuario.
    var user models.User

    //* Escaneamos la fila directamente en la estructura user.
    //* QueryRow().Scan() es suficiente para manejar una única fila.
    if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
        //* Si no se encuentra el usuario o ocurre algún error durante el escaneo, lo registramos y retornamos el error.
        log.Printf("Error al escanear el usuario: %v", err)
        return models.User{}, err
    }

    //* Retornamos el usuario encontrado con sus datos.
    return user, nil
}

