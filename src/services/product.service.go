package services

import (
	"database/sql"
	"fmt"
	"log"

	"firstAPI-go/src/models"
)

type ProductService struct {
	DB *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{DB: db}
}

func (s *ProductService) CreateProduct(product models.Product) error {

	query := "INSERT INTO product (name, price, stock, id_category) VALUES (?, ?, ?, ?)"

	_, err := s.DB.Exec(query, product.Name, product.Price, product.Stock, product.ID_category)
	if err != nil {
		log.Printf("Error al crear el producto: %v", err)
		return err
	}

	fmt.Println("Producto creado con éxito!")

	return nil

}
