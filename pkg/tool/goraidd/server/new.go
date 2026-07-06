package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
)

func New(
	s *store.Store,
	outputPath string,
	r face.Reporter,
) *Server {
	return &Server{
		store:      s,
		outputPath: outputPath,
		reporter:   r,
	}
}
