package controller

import (
	"errors"

	gql "github.com/graphql-go/graphql"
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
		return models.Product{}, errors.New("invalid input")
	}
	productInput := models.ProductInput{
		Name:        input["name"].(string),
		Description: input["description"].(string),
		Price:       input["price"].(float64),
	}
	msgs, isValid := helpers.Validate(productInput)
	if !isValid {
		return models.Product{}, errors.New(msgs[0])
	}
	return p.usecase.CreateProduct(productInput)
}

// DeleteProduct implements [ProductControllerInterface].
func (p *ProductControllerGraphQL) DeleteProduct(param gql.ResolveParams) (any, error) {
	id, ok := param.Args["id"].(int)
	if !ok {
		return models.Product{}, errors.New("invalid id")
	}
	return p.usecase.DeleteProduct(id)
}

// GetProductById implements [ProductControllerInterface].
func (p *ProductControllerGraphQL) GetProductById(param gql.ResolveParams) (any, error) {
	id, ok := param.Args["id"].(int)
	if !ok {
		return models.Product{}, errors.New("invalid id")
	}
	return p.usecase.GetProductById(id)
}

// GetProducts implements [ProductControllerInterface].
func (p *ProductControllerGraphQL) GetProducts(param gql.ResolveParams) (any, error) {
	return p.usecase.GetProducts()
}

// UpdateProduct implements [ProductControllerInterface].
func (p *ProductControllerGraphQL) UpdateProduct(param gql.ResolveParams) (any, error) {
	id, ok := param.Args["id"].(int)
	if !ok {
		return models.Product{}, errors.New("invalid id")
	}
	input, ok := param.Args["input"].(map[string]any)
	if !ok {
		return models.Product{}, errors.New("invalid input")
	}
	var productInput models.ProductUpdateInput
	if name, exists := input["name"]; exists {
		nameStr, ok := name.(string)
		if !ok {
			return models.Product{}, errors.New("invalid name")
		}
		productInput.Name = &nameStr
	}
	if description, exists := input["description"]; exists {
		descriptionStr, ok := description.(string)
		if !ok {
			return models.Product{}, errors.New("invalid description")
		}
		productInput.Description = &descriptionStr
	}
	if price, exists := input["price"]; exists {
		priceFloat, ok := price.(float64)
		if !ok {
			return models.Product{}, errors.New("invalid price")
		}
		productInput.Price = &priceFloat
	}

	msgs, isValid := helpers.Validate(productInput)
	if !isValid {
		return models.Product{}, errors.New(msgs[0])
	}
	return p.usecase.UpdateProduct(id, productInput)
}
