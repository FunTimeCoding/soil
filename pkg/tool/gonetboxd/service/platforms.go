package service

import "github.com/funtimecoding/soil/pkg/netbox/platform"

func (s *Service) Platforms() ([]*platform.Platform, error) {
	return guarded(
		s.guard,
		func() ([]*platform.Platform, error) {
			return s.client.Platforms()
		},
	)
}
