package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	secret "github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
)

func New(
	v secret.Secrets,
	r face.Reporter,
) *Server {
	return &Server{service: v, reporter: r}
}
