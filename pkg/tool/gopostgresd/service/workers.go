package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
)

func (s *Service) Workers() []face.Worker {
	var result []face.Worker

	for _, i := range s.inventory.Instances {
		p, e := s.pool.Client(i.Name)
		errors.PanicOnError(e)
		result = append(result, p.Guard())
	}

	return result
}
