package service

import "github.com/funtimecoding/soil/pkg/linkace/list"

func (s *Service) List(name string) (*list.List, error) {
	return guarded(
		s.guard,
		func() (*list.List, error) { return s.client.List(name) },
	)
}
