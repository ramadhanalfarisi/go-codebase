package repository

import "github.com/ramadhanalfarisi/go-codebase/services/product/models"

type ProductRepositoryInterface interface {
	GetProducts() ([]models.Product, error)
	GetProductById(id int) (models.Product, error)
	CreateProduct(input models.CreateProductInput) (models.Product, error)
	UpdateProduct(id int, input models.PatchProductInput) (models.Product, error)
	UpdatePutProduct(id int, input models.PutProductInput) (models.Product, error)
	DeleteProduct(id int) error
}