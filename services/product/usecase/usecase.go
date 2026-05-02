package usecase

import (
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
func (p *ProductUsecase) CreateProduct(input models.CreateProductInput) (models.Product, error) {
	prod, err := p.repository.CreateProduct(input)
	if err != nil {
		return models.Product{}, errors.New("failed to create product")
	}
	return prod, nil
}

// DeleteProduct implements [ProductUsecaseInterface].
func (p *ProductUsecase) DeleteProduct(id int) (models.Product, error) {
	prod, err := p.repository.GetProductById(id)
	if err != nil {
		return models.Product{}, errors.New("product not found")
	}

	err = p.repository.DeleteProduct(id)
	if err != nil {
		return models.Product{}, errors.New("failed to delete product")
	}
	return prod, nil
}

// GetProductById implements [ProductUsecaseInterface].
func (p *ProductUsecase) GetProductById(id int) (models.Product, error) {
	prod, err := p.repository.GetProductById(id)
	if err != nil {
		return models.Product{}, errors.New("failed to get product")
	}
	return prod, nil
}

// GetProducts implements [ProductUsecaseInterface].
func (p *ProductUsecase) GetProducts() ([]models.Product, error) {
	products, err := p.repository.GetProducts()
	if err != nil {
		return nil, errors.New("failed to get products")
	}
	return products, nil
}

// UpdateProduct implements [ProductUsecaseInterface].
func (p *ProductUsecase) UpdateProduct(id int, input models.PatchProductInput) (models.Product, error) {
	prod, err := p.repository.UpdateProduct(id, input)
	if err != nil {
		return models.Product{}, errors.New("failed to update product")
	}
	return prod, nil
}

// UpdatePutProduct implements [ProductUsecaseInterface].
func (p *ProductUsecase) UpdatePutProduct(id int, input models.PutProductInput) (models.Product, error) {
	prod, err := p.repository.UpdatePutProduct(id, input)
	if err != nil {
		return models.Product{}, errors.New("failed to update product")
	}
	return prod, nil
}
