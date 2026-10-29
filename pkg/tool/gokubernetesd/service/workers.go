package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"maps"
	"slices"
)

func (s *Service) Workers() []face.Worker {
	var result []face.Worker

	for _, name := range slices.Sorted(maps.Keys(s.guards)) {
		result = append(result, s.guards[name])
	}

	return result
}
