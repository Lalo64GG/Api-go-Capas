package services

import (
	"database/sql"
	"firstAPI-go/src/models"
	"fmt"
	"log"
)

type CategoryService struct {
	DB *sql.DB
}

func NewCategoryService(db *sql.DB) *CategoryService {
	return &CategoryService{DB: db}
}

func (s *CategoryService) CreateCategory(category models.Category) error {
	query :=  "INSERT INTO category (name) VALUES (?)"

	_, err := s.DB.Exec(query, category.Name)
	if err != nil {
		log.Printf("Error al crear la categoría: %v", err)
		return err
	}

	fmt.Print("Categoría creadado con exito!")

	return nil
}