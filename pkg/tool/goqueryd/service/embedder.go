package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) Embedder() face.ReachableEmbedder {
	return s.embedder
}
