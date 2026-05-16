package graphql

import (
	"database/sql"

	gql "github.com/graphql-go/graphql"
	"github.com/ramadhanalfarisi/go-codebase/services/common/models"
	"github.com/ramadhanalfarisi/go-codebase/services/product/routes"
)

type Root struct {
	Query    *gql.Object
	Mutation *gql.Object
}

func NewRoot(db *sql.DB) *Root {
	queryFields, mutationFields := initRoutes(db)
	queryObj := gql.NewObject(gql.ObjectConfig{
		Name:   "Query",
		Fields: queryFields,
	})

	mutationObj := gql.NewObject(gql.ObjectConfig{
		Name:   "Mutation",
		Fields: mutationFields,
	})

	return &Root{Query: queryObj, Mutation: mutationObj}
}

func initRoutes(db *sql.DB) (queryFields gql.Fields, mutationFields gql.Fields) {
	querySlices, mutationSlices := loadAllRoutes(db)
	queryFields = gql.Fields{}
	mutationFields = gql.Fields{}
	for _, querySlice := range querySlices {
		for _, query := range querySlice {
			queryFields[query.Name] = &gql.Field{
				Name:    query.Name,
				Type:    query.Response,
				Resolve: query.Controller,
			}
			if query.Request != nil {
				queryFields[query.Name].Args = gql.FieldConfigArgument{
					"filter": &gql.ArgumentConfig{
						Type: gql.NewNonNull(query.Request),
					},
				}
			}
		}
	}
	for _, mutationSlice := range mutationSlices {
		for _, mutation := range mutationSlice {
			mutationFields[mutation.Name] = &gql.Field{
				Name:    mutation.Name,
				Type:    mutation.Response,
				Resolve: mutation.Controller,
			}
			if mutation.Request != nil {
				mutationFields[mutation.Name].Args = gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(mutation.Request),
					},
				}
			}
		}
	}
	return queryFields, mutationFields
}

func loadAllRoutes(db *sql.DB) (querySlices [][]models.GraphQLObjectModel, mutationSlices [][]models.GraphQLObjectModel) {
	productQuery, productMutation := routes.ProductGraphQLRoutes(db)
	querySlices = [][]models.GraphQLObjectModel{
		productQuery,
	}
	mutationSlices = [][]models.GraphQLObjectModel{
		productMutation,
	}
	return
}

