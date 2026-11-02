package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) Workers() []face.Worker {
	result := []face.Worker{s.reaper}

	for _, g := range s.guards {
		result = append(result, g)
	}

	return append(result, s.runner)
}
