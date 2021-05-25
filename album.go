package main

import "github.com/graphql-go/graphql"

var _ AlbumService = (*albumService)(nil)

type AlbumService interface {
	Albums(p graphql.ResolveParams) (interface{}, error)
	Album(p graphql.ResolveParams) (interface{}, error)
}

var albumType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Album",
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
		"photos": &graphql.Field{
			Type: graphql.NewList(photoType),
		},
	}},
)

type albumService struct{ *Resolver }

func (a *albumService) Albums(p graphql.ResolveParams) (interface{}, error) {
	return a.Client.Album.List(p.Context)
}

func (a *albumService) Album(p graphql.ResolveParams) (interface{}, error) {
	input := uint64(p.Args["id"].(int))
	return a.Client.Album.Get(p.Context, input)
}
