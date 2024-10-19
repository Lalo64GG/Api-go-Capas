package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"firstAPI-go/src/config"
	_ "github.com/go-sql-driver/mysql"
)


func main() {

	db, err := config.InitDB()

	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	if err := Migrate(db); err != nil {
		log.Fatalf("Error ejecutando las migraciones: %v", err)
	}

}


func Migrate (db *sql.DB) error {
	migrationStatements := []string{
		"src/config/migrates/sql/00_create_users_table.sql",
		"src/config/migrates/sql/01_create_category_table.sql",
		"src/config/migrates/sql/02_create_products_table.sql",
	}

  
	for _,file := range migrationStatements {

		err := executeSqlFiles(db, file)
		if err != nil {
			log.Fatalf("Error al ejecutar archivo %s: %v", file, err)
		}

	}

	log.Println("Migraciones aplicadas con éxito")
	return nil
}

func executeSqlFiles(db *sql.DB, filepath string) error {

	content, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("no se pudo leer el archivo %w", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("error al ejecutar el archivo SQL: %w", err)
	}

	return nil

}
