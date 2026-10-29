package service

import "github.com/funtimecoding/soil/pkg/opnsense/forward"

func (s *Service) Forwards(phrase string) ([]*forward.Forward, error) {
	return s.source.Forwards(phrase)
}
