package usecase

import (
	"context"

	"github.com/ramadhanalfarisi/go-codebase/services/product/models"
)

type ProductUsecaseInterface interface {
	GetProducts(ctx context.Context, ) ([]models.Product, error)
	GetProductById(ictx context.Context, d int) (models.Product, error)
	CreateProduct(ctx context.Context, input models.CreateProductInput) (models.Product, error)
	UpdateProduct(ctx context.Context, id int, input models.PatchProductInput) (models.Product, error)
	UpdatePutProduct(ctx context.Context, id int, input models.PutProductInput) (models.Product, error)
	DeleteProduct(ctx context.Context, id int) (models.Product, error)
}
