package base

import "github.com/funtimecoding/soil/pkg/face"

func (s *Server) Embedder() face.ReachableEmbedder {
	return s.embedder
}
