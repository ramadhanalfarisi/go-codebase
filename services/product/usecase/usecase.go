package usecase

import (
	"context"
	"errors"

	"github.com/ramadhanalfarisi/go-codebase/services/product/models"
	"github.com/ramadhanalfarisi/go-codebase/services/product/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepositoryInterface
}

func NewProductUsecase(repository repository.ProductRepositoryInterface) ProductUsecaseInterface {
	return &ProductUsecase{repository: repository}
}

// CreateProduct implements [ProductUsecaseInterface].
func (p *ProductUsecase) CreateProduct(ctx context.Context, input models.CreateProductInput) (models.Product, error) {
	prod, err := p.repository.CreateProduct(ctx, input)
	if err != nil {
		return models.Product{}, errors.New("failed to create product")
	}
	return prod, nil
}

// DeleteProduct implements [ProductUsecaseInterface].
func (p *ProductUsecase) DeleteProduct(ctx context.Context, id int) (models.Product, error) {
	prod, err := p.repository.GetProductById(ctx, id)
	if err != nil {
		return models.Product{}, errors.New("product not found")
	}

	err = p.repository.DeleteProduct(ctx, id)
	if err != nil {
		return models.Product{}, errors.New("failed to delete product")
	}
	return prod, nil
}

// GetProductById implements [ProductUsecaseInterface].
func (p *ProductUsecase) GetProductById(ctx context.Context, id int) (models.Product, error) {
	prod, err := p.repository.GetProductById(ctx, id)
	if err != nil {
		return models.Product{}, errors.New("failed to get product")
	}
	return prod, nil
}

// GetProducts implements [ProductUsecaseInterface].
func (p *ProductUsecase) GetProducts(ctx context.Context) ([]models.Product, error) {
	products, err := p.repository.GetProducts(ctx)
	if err != nil {
		return nil, errors.New("failed to get products")
	}
	return products, nil
}

// UpdateProduct implements [ProductUsecaseInterface].
func (p *ProductUsecase) UpdateProduct(ctx context.Context, id int, input models.PatchProductInput) (models.Product, error) {
	prod, err := p.repository.UpdateProduct(ctx, id, input)
	if err != nil {
		return models.Product{}, errors.New("failed to update product")
	}
	return prod, nil
}

// UpdatePutProduct implements [ProductUsecaseInterface].
func (p *ProductUsecase) UpdatePutProduct(ctx context.Context, id int, input models.PutProductInput) (models.Product, error) {
	prod, err := p.repository.UpdatePutProduct(ctx, id, input)
	if err != nil {
		return models.Product{}, errors.New("failed to update product")
	}
	return prod, nil
}
