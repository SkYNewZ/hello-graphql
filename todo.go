package main

import "github.com/graphql-go/graphql"

var _ TodoService = (*todoService)(nil)

type TodoService interface {
	Todos(p graphql.ResolveParams) (interface{}, error)
	Todo(p graphql.ResolveParams) (interface{}, error)
}

var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"userId": &graphql.Field{
			Type: graphql.Int,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"completed": &graphql.Field{
			Type: graphql.Boolean,
		},
	}},
)

type todoService struct{ *Resolver }

func (t *todoService) Todos(p graphql.ResolveParams) (interface{}, error) {
	return t.Client.Todo.List(p.Context)
}

func (t *todoService) Todo(p graphql.ResolveParams) (interface{}, error) {
	input := uint64(p.Args["id"].(int))
	return t.Client.Todo.Get(p.Context, input)
}
