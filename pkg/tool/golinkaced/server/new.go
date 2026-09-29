package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	linkace "github.com/funtimecoding/soil/pkg/tool/golinkaced/face"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/service"
)

func New(
	c linkace.LinkAceSource,
	s *service.Service,
	r face.Reporter,
) *Server {
	return &Server{client: c, service: s, reporter: r}
}
