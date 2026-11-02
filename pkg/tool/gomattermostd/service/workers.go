package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) Workers() []face.Worker {
	result := []face.Worker{s.guard}

	if s.watcher != nil {
		result = append(result, s.rosterGuard, s.watcher)
	}

	return result
}
