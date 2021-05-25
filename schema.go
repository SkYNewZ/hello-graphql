package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func newGraphQLHandler() (http.Handler, error) {
	resolver := NewResolver()
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    resolver.Query(),
		Mutation: resolver.Mutation(),
	})
	if err != nil {
		return nil, err
	}

	return gqlhandler.New(&gqlhandler.Config{
		Schema:     &schema,
		Playground: true,
	}), nil
}
