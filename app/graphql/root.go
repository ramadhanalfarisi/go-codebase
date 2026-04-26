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
	productQuery, productMutation := routes.ProductRoutes(db)
	querySlices := [][]models.GraphQLObjectModel{
		productQuery,
	}
	mutationSlices := [][]models.GraphQLObjectModel{
		productMutation,
	}
	queryFields = gql.Fields{}
	mutationFields = gql.Fields{}
	for _, querySlice := range querySlices {
		for _, query := range querySlice {
			queryFields[query.Name] = query.Field
		}
	}
	for _, mutationSlice := range mutationSlices {
		for _, mutation := range mutationSlice {
			mutationFields[mutation.Name] = mutation.Field
		}
	}
	return queryFields, mutationFields
}
