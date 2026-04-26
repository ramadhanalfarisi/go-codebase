package controller

import (
	gql "github.com/graphql-go/graphql"
)

type ProductControllerInterface interface {
	GetProducts(param gql.ResolveParams) (any, error)
	GetProductById(param gql.ResolveParams) (any, error)
	CreateProduct(param gql.ResolveParams) (any, error)
	UpdateProduct(param gql.ResolveParams) (any, error)
	DeleteProduct(param gql.ResolveParams) (any, error)
}
