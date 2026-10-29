package service

import "github.com/funtimecoding/soil/pkg/errors/validation"

func (s *Service) Click(label string) error {
	if !s.labels[label] {
		return validation.New("unknown label: %s", label)
	}

	return s.store.Create(label)
}
