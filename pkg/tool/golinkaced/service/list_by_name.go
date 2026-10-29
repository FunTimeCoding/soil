package service

import "github.com/funtimecoding/soil/pkg/linkace/list"

func (s *Service) ListByName(name string) (*list.List, error) {
	return guarded(
		s.guard,
		func() (*list.List, error) { return s.client.ListByName(name) },
	)
}
