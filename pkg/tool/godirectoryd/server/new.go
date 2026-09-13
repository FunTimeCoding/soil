package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
)

func New(
	v *service.Service,
	r face.Reporter,
) *Server {
	return &Server{service: v, reporter: r}
}
