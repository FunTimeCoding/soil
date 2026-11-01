package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) Stopper() face.Worker {
	return s.stop
}
