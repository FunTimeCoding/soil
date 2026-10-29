package service

import "github.com/funtimecoding/soil/pkg/linkace/list"

func (s *Service) UpdateList(
	identifier int,
	patch map[string]any,
) (*list.List, error) {
	return guarded(
		s.guard,
		func() (*list.List, error) {
			return s.client.UpdateList(identifier, patch)
		},
	)
}
