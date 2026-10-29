package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) Workers() []face.Worker {
	if s.directoryGuard == nil {
		return nil
	}

	return []face.Worker{s.directoryGuard}
}
