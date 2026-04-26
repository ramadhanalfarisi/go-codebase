package repository

import "github.com/ramadhanalfarisi/go-codebase/services/product/models"

type ProductRepositoryInterface interface {
	GetProducts() ([]models.Product, error)
	GetProductById(id int) (models.Product, error)
	CreateProduct(input models.ProductInput) (models.Product, error)
	UpdateProduct(id int, input models.ProductUpdateInput) (models.Product, error)
	DeleteProduct(id int) error
}