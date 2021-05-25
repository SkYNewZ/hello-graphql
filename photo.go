package main

import "github.com/graphql-go/graphql"

var _ PhotoService = (*photoService)(nil)

type PhotoService interface {
	Photos(p graphql.ResolveParams) (interface{}, error)
	Photo(p graphql.ResolveParams) (interface{}, error)
}

var photoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Photo",
	Fields: graphql.Fields{
		"albumId": &graphql.Field{
			Type: graphql.Int,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"url": &graphql.Field{
			Type: graphql.String,
		},
		"thumbnailUrl": &graphql.Field{
			Type: graphql.String,
		},
	}},
)

type photoService struct{ *Resolver }

func (p *photoService) Photos(params graphql.ResolveParams) (interface{}, error) {
	return p.Client.Photo.List(params.Context)
}

func (p *photoService) Photo(params graphql.ResolveParams) (interface{}, error) {
	input := uint64(params.Args["id"].(int))
	return p.Client.Photo.Get(params.Context, input)
}
