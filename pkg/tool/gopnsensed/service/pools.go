package service

import "github.com/funtimecoding/soil/pkg/opnsense/pool"

func (s *Service) Pools(phrase string) ([]*pool.Pool, error) {
	return guarded(
		s.guard,
		func() ([]*pool.Pool, error) {
			return s.client.Pools(phrase)
		},
	)
}
