package graphql

import (
	"fmt"
	"net/http"

	gql "github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ramadhanalfarisi/go-codebase/config"
	"github.com/ramadhanalfarisi/go-codebase/drivers"
)

type GraphQL struct {
	handler *handler.Handler
}

func NewGraphQL() *GraphQL {
	db := drivers.ConnectDB()
	object := NewRoot(db)
	schema, _ := gql.NewSchema(gql.SchemaConfig{
		Query:    object.Query,
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
	fmt.Println("Your application running on http://localhost:" + config.PORT_GRAPHQL)
	http.Handle("/graphql", g.handler)
	http.ListenAndServe(config.PORT_GRAPHQL, nil)
}
