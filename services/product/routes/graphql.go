package routes

import (
	"database/sql"

	gql "github.com/graphql-go/graphql"
	"github.com/ramadhanalfarisi/go-codebase/services/common/models"
	"github.com/ramadhanalfarisi/go-codebase/services/product/controller"
	"github.com/ramadhanalfarisi/go-codebase/services/product/repository"
	"github.com/ramadhanalfarisi/go-codebase/services/product/usecase"
)

var productType = gql.NewObject(gql.ObjectConfig{
	Name: "Product",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type: gql.Int,
		},
		"name": &gql.Field{
			Type: gql.String,
		},
		"description": &gql.Field{
			Type: gql.String,
		},
		"price": &gql.Field{
			Type: gql.Float,
		},
	},
})

func ProductGraphQLRoutes(db *sql.DB) (productQuery []models.GraphQLObjectModel, productMutation []models.GraphQLObjectModel) {
	repo := repository.NewProductRepository(db)
	usecase := usecase.NewProductUsecase(repo)
	controller := controller.NewProductControllerGraphQL(usecase)

	productQuery = []models.GraphQLObjectModel{
	{
		Name: "Products",
		Field: &gql.Field{
			Type: gql.NewList(productType),
			Resolve: controller.GetProducts,
		},
	},
	{
		Name: "Product",
		Field: &gql.Field{
			Type: gql.NewList(productType),
			Args: gql.FieldConfigArgument{
				"id": &gql.ArgumentConfig{
					Type: gql.Int,
				},
			},
			Resolve: controller.GetProductById,
		},
	}}

	productMutation = []models.GraphQLObjectModel{
	{
		Name: "CreateProduct",
		Field: &gql.Field{
			Type: productType,
			Args: gql.FieldConfigArgument{
				"name": &gql.ArgumentConfig{
					Type: gql.String,
				},
				"description": &gql.ArgumentConfig{
					Type: gql.String,
				},
				"price": &gql.ArgumentConfig{
					Type: gql.Float,
				},
			},
			Resolve: controller.CreateProduct,
		}},
	{
		Name: "UpdateProduct",
		Field: &gql.Field{
			Type: productType,
			Args: gql.FieldConfigArgument{
				"id": &gql.ArgumentConfig{
					Type: gql.Int,
				},
				"name": &gql.ArgumentConfig{
					Type: gql.String,
				},
				"description": &gql.ArgumentConfig{
					Type: gql.String,
				},
				"price": &gql.ArgumentConfig{
					Type: gql.Float,
				},
			},
			Resolve: controller.UpdateProduct,
		}}, {
		Name: "DeleteProduct",
		Field: &gql.Field{
			Type: productType,
			Args: gql.FieldConfigArgument{
				"id": &gql.ArgumentConfig{
					Type: gql.Int,
				},
			},
			Resolve: controller.DeleteProduct,
		},
	}}
	return productQuery, productMutation
}
