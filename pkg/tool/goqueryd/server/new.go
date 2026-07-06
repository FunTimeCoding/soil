package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
)

func New(
	s *service.Service,
	r face.Reporter,
) *Server {
	return &Server{
		service:  s,
		reporter: r,
	}
}
