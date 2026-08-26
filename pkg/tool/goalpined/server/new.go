package server

import "github.com/funtimecoding/soil/pkg/face"

func New(r face.Reporter) *Server {
	return &Server{reporter: r}
}
