package controller

import (
	"errors"

	gql "github.com/graphql-go/graphql"
	"github.com/ramadhanalfarisi/go-codebase/constants"
	"github.com/ramadhanalfarisi/go-codebase/helpers"
	"github.com/ramadhanalfarisi/go-codebase/services/product/models"
	"github.com/ramadhanalfarisi/go-codebase/services/product/usecase"
)

type ProductControllerGraphQL struct {
	usecase usecase.ProductUsecaseInterface
}

func NewProductControllerGraphQL(usecase usecase.ProductUsecaseInterface) ProductControllerGraphQLInterface {
	return &ProductControllerGraphQL{usecase: usecase}
}

// CreateProduct implements [ProductControllerInterface].
func (p *ProductControllerGraphQL) CreateProduct(param gql.ResolveParams) (any, error) {
	input, ok := param.Args["input"].(map[string]any)
	if !ok {
		return models.Product{}, errors.New(constants.InvalidInput)
	}
	var productInput models.CreateProductInput
	helpers.CollectGraphqlArguments(input, &productInput)

	msgs, isValid := helpers.Validate(productInput)
	if !isValid {
		return models.Product{}, errors.New(msgs[0])
	}
	return p.usecase.CreateProduct(param.Context, productInput)
}

// DeleteProduct implements [ProductControllerInterface].
func (p *ProductControllerGraphQL) DeleteProduct(param gql.ResolveParams) (any, error) {
	input, ok := param.Args["input"].(map[string]any)
	if !ok {
		return models.Product{}, errors.New(constants.InvalidInput)
	}
	var productInput models.DeleteProductInput
	helpers.CollectGraphqlArguments(input, &productInput)

	msgs, isValid := helpers.Validate(productInput)
	if !isValid {
		return models.Product{}, errors.New(msgs[0])
	}
	return p.usecase.DeleteProduct(param.Context, productInput.Id)
}

// GetProductById implements [ProductControllerInterface].
func (p *ProductControllerGraphQL) GetProductById(param gql.ResolveParams) (any, error) {
	filter, ok := param.Args["filter"].(map[string]any)
	if !ok {
		return models.Product{}, errors.New(constants.InvalidFilter)
	}
	var productFilter models.ProductFilter
	helpers.CollectGraphqlArguments(filter, &productFilter)

	msgs, isValid := helpers.Validate(productFilter)
	if !isValid {
		return models.Product{}, errors.New(msgs[0])
	}
	return p.usecase.GetProductById(param.Context, productFilter.Id)
}

// GetProducts implements [ProductControllerInterface].
func (p *ProductControllerGraphQL) GetProducts(param gql.ResolveParams) (any, error) {
	return p.usecase.GetProducts(param.Context)
}

// UpdateProduct implements [ProductControllerInterface].
func (p *ProductControllerGraphQL) UpdateProduct(param gql.ResolveParams) (any, error) {
	input, ok := param.Args["input"].(map[string]any)
	if !ok {
		return models.Product{}, errors.New(constants.InvalidInput)
	}
	var productInput models.PatchProductInputGraphql
	helpers.CollectGraphqlArguments(input, &productInput)

	msgs, isValid := helpers.Validate(productInput)
	if !isValid {
		return models.Product{}, errors.New(msgs[0])
	}

	return p.usecase.UpdateProduct(param.Context, productInput.Id, models.PatchProductInput{Name: productInput.Name, Price: productInput.Price, Description: productInput.Description})
}
