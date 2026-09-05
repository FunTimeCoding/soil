package base

import "github.com/funtimecoding/soil/pkg/face"

func (s *Server) Embedder() face.Embedder {
	return s.embedder
}
