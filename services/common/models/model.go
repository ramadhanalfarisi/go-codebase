package models

import (
	"database/sql"

	gql "github.com/graphql-go/graphql"
)

type Model struct {
	DB    *sql.DB
	Model interface{}
	Args  []any
}

type GraphQLObjectModel struct {
	Name string
	Field *gql.Field
}