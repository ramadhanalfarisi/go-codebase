package repository

import (
	"database/sql"
	"time"

	"github.com/ramadhanalfarisi/go-codebase/helpers"
	"github.com/ramadhanalfarisi/go-codebase/helpers/query_builder"
	"github.com/ramadhanalfarisi/go-codebase/services/product/models"
)

type ProductRepository struct {
	db          *sql.DB
	queryHelper helpers.QueryHelperInterface
}

func NewProductRepository(db *sql.DB) ProductRepositoryInterface {
	return &ProductRepository{db: db, queryHelper: helpers.NewQueryHelper(db)}
}

// CreateProduct implements [ProductRepositoryInterface].
func (p *ProductRepository) CreateProduct(input models.CreateProductInput) (models.Product, error) {
	query, args := query_builder.New("products").Insert("name", "description", "price", "created_at").Values(input.Name, input.Description, input.Price, time.Now()).Build("id", "name", "description", "price")
	var product models.Product
	err := p.queryHelper.Insert(query, args, &product.Id, &product.Name, &product.Description, &product.Price)
	if err != nil {
		helpers.Error(err)
		return models.Product{}, err
	}
	return product, nil
}

// DeleteProduct implements [ProductRepositoryInterface].
func (p *ProductRepository) DeleteProduct(id int) error {
	query, args := query_builder.New("products").Delete().Where("id = ?", id).Build()
	err := p.queryHelper.Delete(query, args)
	if err != nil {
		helpers.Error(err)
	}
	return err
}

// GetProductById implements [ProductRepositoryInterface].
func (p *ProductRepository) GetProductById(id int) (models.Product, error) {
	query, args := query_builder.New("products").Select("id", "name", "description", "price").Where("id = ?", id).Build()
	var product models.Product
	err := p.queryHelper.SelectRow(query, args, &product.Id, &product.Name, &product.Description, &product.Price)
	if err != nil {
		helpers.Error(err)
		return models.Product{}, err
	}
	return product, nil
}

// GetProducts implements [ProductRepositoryInterface].
func (p *ProductRepository) GetProducts() ([]models.Product, error) {
	query, args := query_builder.New("products").Select("id", "name", "description", "price").Build()
	var products []models.Product
	err := p.queryHelper.Select(query, args, func() {
		product := models.Product{}
		products = append(products, product)
	})
	if err != nil {
		helpers.Error(err)
		return nil, err
	}
	return products, nil
}

// UpdateProduct implements [ProductRepositoryInterface].
func (p *ProductRepository) UpdateProduct(id int, input models.PatchProductInput) (models.Product, error) {
	update := query_builder.New("products").Update()
	if input.Name != nil {
		update.Set("name", *input.Name)
	}
	if input.Description != nil {
		update.Set("description", *input.Description)
	}
	if input.Price != nil {
		update.Set("price", *input.Price)
	}
	query, args := update.Where("id = ?", id).Build("id", "name", "description", "price")
	var product models.Product
	err := p.queryHelper.Update(query, args, &product.Id, &product.Name, &product.Description, &product.Price)
	if err != nil {
		helpers.Error(err)
	}
	return product, err
}

// UpdatePutProduct implements [ProductRepositoryInterface].
func (p *ProductRepository) UpdatePutProduct(id int, input models.PutProductInput) (models.Product, error) {
	query, args := query_builder.New("products").Update().Set("name", input.Name).Set("description", input.Description).Set("price", input.Price).Where("id = ?", id).Build("id", "name", "description", "price")
	var product models.Product
	err := p.queryHelper.Update(query, args, &product.Id, &product.Name, &product.Description, &product.Price)
	if err != nil {
		helpers.Error(err)
	}
	return product, err
}
