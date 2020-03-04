package main

import (
	"github.com/NelsonMario/connection"
	"github.com/NelsonMario/graphql/mutations"
	"github.com/NelsonMario/graphql/query"
	"github.com/NelsonMario/middleware"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main() {

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query.GetRoot(),
		Mutation: mutations.GetRoot(),
	})

	if err != nil {
		panic(err.Error())
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     false,
		GraphiQL:   true,
		Playground: true,
	})

	m := middleware.Cors(h)
	r := connection.NewRoutes()
	r.Handle("/api/{key}", m)

	log.Fatal(http.ListenAndServe(":"+connection.Port, r))
}

