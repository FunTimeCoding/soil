package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

func New(
	s *store.Store,
	r face.Reporter,
) *Server {
	return &Server{
		store:    s,
		reporter: r,
	}
}
