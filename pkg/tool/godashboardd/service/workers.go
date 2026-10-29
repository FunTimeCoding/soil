package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) Workers() []face.Worker {
	result := []face.Worker{s.prometheusGuard}

	if s.argocdGuard != nil {
		result = append(result, s.argocdGuard)
	}

	if s.usageGuard != nil {
		result = append(result, s.usageGuard)
	}

	return result
}
