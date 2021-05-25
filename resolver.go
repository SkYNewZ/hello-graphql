package main

import "github.com/SkYNewZ/jsonplaceholder"

// Resolver describes our global GraphQL resolver instance
type Resolver struct {
	Client  *jsonplaceholder.Client
	Post    PostService
	Comment CommentService
	Album   AlbumService
	Photo   PhotoService
	Todo    TodoService
	User    UserService
}

func NewResolver() *Resolver {
	resolver := new(Resolver)
	resolver.Client = jsonplaceholder.New(jsonplaceholder.WithHTTPClient(newHTTPClient()))

	resolver.Post = &postService{resolver}
	resolver.Comment = &commentService{resolver}
	resolver.Album = &albumService{resolver}
	resolver.Photo = &photoService{resolver}
	resolver.Todo = &todoService{resolver}
	resolver.User = &userService{resolver}

	return resolver
}
