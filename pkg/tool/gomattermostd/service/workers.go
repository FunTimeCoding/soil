package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) Workers() []face.Worker {
	result := []face.Worker{s.guard}

	if s.watcher != nil {
		result = append(result, s.watcher, s.sweeper)
	}

	if s.monitor != nil {
		result = append(result, s.monitor)
	}

	return result
}
