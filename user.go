package main

import "github.com/graphql-go/graphql"

var _ UserService = (*userService)(nil)

type UserService interface {
	Users(p graphql.ResolveParams) (interface{}, error)
	User(p graphql.ResolveParams) (interface{}, error)
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "address",
				Fields: graphql.Fields{
					"street": &graphql.Field{
						Type: graphql.String,
					},
					"suite": &graphql.Field{
						Type: graphql.String,
					},
					"city": &graphql.Field{
						Type: graphql.String,
					},
					"zipcode": &graphql.Field{
						Type: graphql.String,
					},
					"geo": &graphql.Field{
						Type: graphql.NewObject(graphql.ObjectConfig{
							Name: "geo",
							Fields: graphql.Fields{
								"lat": &graphql.Field{
									Type: graphql.String,
								},
								"lng": &graphql.Field{
									Type: graphql.String,
								},
							},
						}),
					},
				},
			}),
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"website": &graphql.Field{
			Type: graphql.String,
		},
		"company": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "company",
				Fields: graphql.Fields{
					"name": &graphql.Field{
						Type: graphql.String,
					},
					"catchPhrase": &graphql.Field{
						Type: graphql.String,
					},
					"bs": &graphql.Field{
						Type: graphql.String,
					},
				},
			}),
		},
		"albums": &graphql.Field{
			Type: graphql.NewList(albumType),
		},
		"todos": &graphql.Field{
			Type: graphql.NewList(todoType),
		},
		"posts": &graphql.Field{
			Type: graphql.NewList(postType),
		},
	}},
)

type userService struct{ *Resolver }

func (u *userService) Users(p graphql.ResolveParams) (interface{}, error) {
	return u.Client.User.List(p.Context)
}

func (u *userService) User(p graphql.ResolveParams) (interface{}, error) {
	input := uint64(p.Args["id"].(int))
	return u.Client.User.Get(p.Context, input)
}
