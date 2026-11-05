package service

import "github.com/funtimecoding/soil/pkg/opnsense/forward"

func (s *Service) Forwards(phrase string) ([]*forward.Forward, error) {
	return guarded(
		s.guard,
		func() ([]*forward.Forward, error) {
			return s.client.Forwards(phrase)
		},
	)
}
