package graphql

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"

	gql "github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ramadhanalfarisi/go-codebase/config"
	"github.com/ramadhanalfarisi/go-codebase/drivers"
	"github.com/ramadhanalfarisi/go-codebase/middlewares"
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
		Playground: true,
	})
	return &GraphQL{handler: h}
}

func (g *GraphQL) Run() {
	fmt.Println("Your application running on http://localhost" + config.PORT_GRAPHQL)
	go func(){
		log.Println(http.ListenAndServe(config.PPROF_GRAPHQL_PORT, nil))
	}()
	chain := middlewares.Chain(
		middlewares.Recovery,  // outermost: catch panics
		middlewares.Logger,    // log method, path, duration
		middlewares.CORS("*"), // CORS headers
		middlewares.Auth, // Auth validation
	)
	http.Handle("/graphql", chain(g.handler))
	http.ListenAndServe(config.PORT_GRAPHQL, nil)
}
