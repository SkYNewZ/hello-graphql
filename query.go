package main

import "github.com/graphql-go/graphql"

func (r *Resolver) Query() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"posts": &graphql.Field{
				Type:        graphql.NewList(postType),
				Resolve:     r.Post.Posts,
				Description: "Return all posts",
			},
			"post": &graphql.Field{
				Type:        postType,
				Resolve:     r.Post.Post,
				Description: "Return post matching ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "",
					},
				},
			},
			"comments": &graphql.Field{
				Type:        graphql.NewList(commentType),
				Resolve:     r.Comment.Comments,
				Description: "Return all comments",
			},
			"comment": &graphql.Field{
				Type:        commentType,
				Resolve:     r.Comment.Comment,
				Description: "Return comment matching ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "",
					},
				},
			},
			"albums": &graphql.Field{
				Type:        graphql.NewList(albumType),
				Resolve:     r.Album.Albums,
				Description: "Return all albums",
			},
			"album": &graphql.Field{
				Type:        albumType,
				Resolve:     r.Album.Album,
				Description: "Return album matching ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "",
					},
				},
			},
			"photos": &graphql.Field{
				Type:        graphql.NewList(photoType),
				Resolve:     r.Photo.Photos,
				Description: "Return all photos",
			},
			"photo": &graphql.Field{
				Type:        photoType,
				Resolve:     r.Photo.Photo,
				Description: "Return photo matching ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "",
					},
				},
			},
			"todos": &graphql.Field{
				Type:        graphql.NewList(todoType),
				Resolve:     r.Todo.Todos,
				Description: "Return all todos",
			},
			"todo": &graphql.Field{
				Type:        todoType,
				Resolve:     r.Todo.Todo,
				Description: "Return todo matching ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "",
					},
				},
			},
			"users": &graphql.Field{
				Type:        graphql.NewList(userType),
				Resolve:     r.User.Users,
				Description: "Return all user",
			},
			"user": &graphql.Field{
				Type:        userType,
				Resolve:     r.User.User,
				Description: "Return user matching ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "",
					},
				},
			},
		},
	})
}
