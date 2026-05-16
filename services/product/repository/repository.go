package repository

import (
	"context"
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
func (p *ProductRepository) CreateProduct(ctx context.Context, input models.CreateProductInput) (models.Product, error) {
	userId := helpers.GetUserIdFromGraphql(ctx)
	query, args := query_builder.New("products").Insert("user_id", "name", "description", "price", "created_at").Values(userId, input.Name, input.Description, input.Price, time.Now()).Build("id", "name", "description", "price")
	var product models.Product
	err := p.queryHelper.Insert(ctx, query, args, &product.Id, &product.Name, &product.Description, &product.Price)
	if err != nil {
		helpers.Error(err)
		return models.Product{}, err
	}
	return product, nil
}

// DeleteProduct implements [ProductRepositoryInterface].
func (p *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	userId := helpers.GetUserIdFromGraphql(ctx)
	query, args := query_builder.New("products").Delete().Where("id = ?", id).Where("user_id = ?", userId).Build()
	err := p.queryHelper.Delete(ctx, query, args)
	if err != nil {
		helpers.Error(err)
	}
	return err
}

// GetProductById implements [ProductRepositoryInterface].
func (p *ProductRepository) GetProductById(ctx context.Context, id int) (models.Product, error) {
	userId := helpers.GetUserIdFromGraphql(ctx)
	query, args := query_builder.New("products").Select("id", "name", "description", "price").Where("id = ?", id).Where("user_id = ?", userId).Build()
	var product models.Product
	err := p.queryHelper.SelectRow(ctx, query, args, &product.Id, &product.Name, &product.Description, &product.Price)
	if err != nil {
		helpers.Error(err)
		return models.Product{}, err
	}
	return product, nil
}

// GetProducts implements [ProductRepositoryInterface].
func (p *ProductRepository) GetProducts(ctx context.Context) ([]models.Product, error) {
	userId := helpers.GetUserIdFromGraphql(ctx)
	query, args := query_builder.New("products").Select("id", "name", "description", "price").Where("user_id = ?", userId).Build()
	var products []models.Product
	err := p.queryHelper.Select(ctx, query, args, func(r *sql.Rows) error {
		var product models.Product
		if err := r.Scan(&product.Id, &product.Name, &product.Description, &product.Price); err != nil {
			return err
		}
		products = append(products, product)
		return nil
	})
	if err != nil {
		helpers.Error(err)
		return nil, err
	}
	return products, nil
}

// UpdateProduct implements [ProductRepositoryInterface].
func (p *ProductRepository) UpdateProduct(ctx context.Context, id int, input models.PatchProductInput) (models.Product, error) {
	userId := helpers.GetUserIdFromGraphql(ctx)
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
	query, args := update.Where("id = ?", id).Where("user_id = ?", userId).Build("id", "name", "description", "price")
	var product models.Product
	err := p.queryHelper.Update(ctx, query, args, &product.Id, &product.Name, &product.Description, &product.Price)
	if err != nil {
		helpers.Error(err)
	}
	return product, err
}

// UpdatePutProduct implements [ProductRepositoryInterface].
func (p *ProductRepository) UpdatePutProduct(ctx context.Context, id int, input models.PutProductInput) (models.Product, error) {
	userId := helpers.GetUserIdFromGraphql(ctx)
	query, args := query_builder.New("products").Update().Set("name", input.Name).Set("description", input.Description).Set("price", input.Price).Where("id = ?", id).Where("user_id = ?", userId).Build("id", "name", "description", "price")
	var product models.Product
	err := p.queryHelper.Update(ctx, query, args, &product.Id, &product.Name, &product.Description, &product.Price)
	if err != nil {
		helpers.Error(err)
	}
	return product, err
}
