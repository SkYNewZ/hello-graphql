package main

import "github.com/graphql-go/graphql"

var _ CommentService = (*commentService)(nil)

type CommentService interface {
	Comments(p graphql.ResolveParams) (interface{}, error)
	Comment(p graphql.ResolveParams) (interface{}, error)
}

var commentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"postId": &graphql.Field{
			Type: graphql.Int,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"body": &graphql.Field{
			Type: graphql.String,
		},
	}},
)

type commentService struct{ *Resolver }

func (c *commentService) Comments(p graphql.ResolveParams) (interface{}, error) {
	return c.Client.Comment.List(p.Context)
}

func (c *commentService) Comment(p graphql.ResolveParams) (interface{}, error) {
	input := uint64(p.Args["id"].(int))
	return c.Client.Comment.Get(p.Context, input)
}
