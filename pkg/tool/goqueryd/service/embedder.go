package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) Embedder() face.Embedder {
	return s.embedder
}
