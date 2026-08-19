package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
)

func New(
	s *store.Store,
	v *service.Service,
	r face.Reporter,
) *Server {
	return &Server{store: s, service: v, reporter: r}
}
