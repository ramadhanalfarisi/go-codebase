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
			Name:       "products",
			Response:   gql.NewList(productType),
			Controller: controller.GetProducts,
		},
		{
			Name:     "product",
			Response: productType,
			Request: gql.NewInputObject(gql.InputObjectConfig{
				Name: "ProductFilter",
				Fields: gql.InputObjectConfigFieldMap{
					"id": &gql.InputObjectFieldConfig{
						Type: gql.Int,
					},
				},
			}),
			Controller: controller.GetProductById,
		},
	}

	productMutation = []models.GraphQLObjectModel{
		{
			Name:     "createProduct",
			Response: productType,
			Request: gql.NewInputObject(gql.InputObjectConfig{
				Name: "CreateProductInput",
				Fields: gql.InputObjectConfigFieldMap{
					"name": &gql.InputObjectFieldConfig{
						Type: gql.String,
					},
					"description": &gql.InputObjectFieldConfig{
						Type: gql.String,
					},
					"price": &gql.InputObjectFieldConfig{
						Type: gql.Float,
					},
				},
			}),
			Controller: controller.CreateProduct,
		},
		{
			Name:     "updateProduct",
			Response: productType,
			Request: gql.NewInputObject(gql.InputObjectConfig{
				Name: "UpdateProductInput",
				Fields: gql.InputObjectConfigFieldMap{
					"id": &gql.InputObjectFieldConfig{
						Type: gql.Int,
					},
					"name": &gql.InputObjectFieldConfig{
						Type: gql.String,
					},
					"description": &gql.InputObjectFieldConfig{
						Type: gql.String,
					},
					"price": &gql.InputObjectFieldConfig{
						Type: gql.Float,
					},
				},
			}),
			Controller: controller.UpdateProduct,
		},
		{
			Name:     "deleteProduct",
			Response: productType,
			Request: gql.NewInputObject(gql.InputObjectConfig{
				Name: "DeleteFilterInput",
				Fields: gql.InputObjectConfigFieldMap{
					"id": &gql.InputObjectFieldConfig{
						Type: gql.Int,
					},
				},
			}),
			Controller: controller.DeleteProduct,
		},
	}
	return productQuery, productMutation
}
