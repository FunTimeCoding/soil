package service

import "github.com/funtimecoding/soil/pkg/opnsense/state"

func (s *Service) States(phrase string) ([]*state.State, error) {
	return s.source.States(phrase)
}
