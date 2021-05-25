package main

import "github.com/graphql-go/graphql"

var _ PostService = (*postService)(nil)

type PostService interface {
	Posts(p graphql.ResolveParams) (interface{}, error)
	Post(p graphql.ResolveParams) (interface{}, error)
}

var postType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
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
		"body": &graphql.Field{
			Type: graphql.String,
		},
		"comments": &graphql.Field{
			Type: graphql.NewList(commentType),
		},
	}},
)

type postService struct{ *Resolver }

func (p *postService) Posts(params graphql.ResolveParams) (interface{}, error) {
	return p.Client.Post.List(params.Context)
}

func (p *postService) Post(params graphql.ResolveParams) (interface{}, error) {
	input := uint64(params.Args["id"].(int))
	return p.Client.Post.Get(params.Context, input)
}
