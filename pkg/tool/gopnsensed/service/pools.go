package service

import "github.com/funtimecoding/soil/pkg/opnsense/pool"

func (s *Service) Pools(phrase string) ([]*pool.Pool, error) {
	return s.source.Pools(phrase)
}
