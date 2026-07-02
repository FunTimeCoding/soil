package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/service"
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
