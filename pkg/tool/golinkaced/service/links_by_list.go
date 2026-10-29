package service

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (s *Service) LinksByList(identifier int) ([]*link.Link, error) {
	return guarded(
		s.guard,
		func() ([]*link.Link, error) { return s.client.LinksByList(identifier) },
	)
}
