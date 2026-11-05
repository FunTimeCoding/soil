package service

import "github.com/funtimecoding/soil/pkg/opnsense/host"

func (s *Service) Hosts(phrase string) ([]*host.Host, error) {
	return guarded(
		s.guard,
		func() ([]*host.Host, error) {
			return s.client.Hosts(phrase)
		},
	)
}
