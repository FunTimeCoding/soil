package service

import "github.com/funtimecoding/soil/pkg/linkace/link"

func (s *Service) updateLink(
	identifier int,
	patch map[string]any,
) (*link.Link, error) {
	return guarded(
		s.guard,
		func() (*link.Link, error) {
			return s.client.UpdateLink(identifier, patch)
		},
	)
}
