package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //* El subrayado indica que solo importamos el paquete para ejecutar su init()
)

var DB *sql.DB

//? InitDB inicializa la conexión a la base de datos y la almacena en una variable global DB
func InitDB() (*sql.DB, error) {
	//* Configuración de la conexión a la base de datos (modifica los parámetros según sea necesario)
	dsn := "root:Lolasso1012@tcp(127.0.0.1:3306)/api_taller"

	//* Abrimos la conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la conexión a la base de datos: %v", err)
	}

	//* Verificamos si la conexión es exitosa
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	} else {
		fmt.Println("Conexión a la base de datos exitosa!")
	}

	DB = db
	return db, nil
}

