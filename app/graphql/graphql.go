package graphql

import (
	"net/http"

	gql "github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ramadhanalfarisi/go-codebase/config"
	"github.com/ramadhanalfarisi/go-codebase/db"
)

type GraphQL struct {
	handler *handler.Handler
}

func NewGraphQL() *GraphQL {
	db := db.ConnectDB()
	object := NewRoot(db)
	schema, _ := gql.NewSchema(gql.SchemaConfig{
		Query: object.Query,
		Mutation: object.Mutation,
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	return &GraphQL{handler: h}
}

func (g *GraphQL) Run() {
	http.Handle("/graphql", g.handler)
	http.ListenAndServe(config.PORT_GRAPHQL, nil)
}
