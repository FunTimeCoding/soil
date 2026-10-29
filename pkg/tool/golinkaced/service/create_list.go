package service

import "github.com/funtimecoding/soil/pkg/linkace/list"

func (s *Service) CreateList(
	name string,
	description string,
) (*list.List, error) {
	return guarded(
		s.guard,
		func() (*list.List, error) {
			return s.client.CreateList(name, description)
		},
	)
}
