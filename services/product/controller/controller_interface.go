package controller

import (
	"github.com/gofiber/fiber/v3"
	gql "github.com/graphql-go/graphql"
)

type ProductControllerGraphQLInterface interface {
	GetProducts(param gql.ResolveParams) (any, error)
	GetProductById(param gql.ResolveParams) (any, error)
	CreateProduct(param gql.ResolveParams) (any, error)
	UpdateProduct(param gql.ResolveParams) (any, error)
	DeleteProduct(param gql.ResolveParams) (any, error)
}

type ProductControllerAPIInterface interface {
	GetProducts(c fiber.Ctx) error
	GetProductById(c fiber.Ctx) error
	CreateProduct(c fiber.Ctx) error
	UpdatePatchProduct(c fiber.Ctx) error
	UpdatePutProduct(c fiber.Ctx) error
	DeleteProduct(c fiber.Ctx) error
}
