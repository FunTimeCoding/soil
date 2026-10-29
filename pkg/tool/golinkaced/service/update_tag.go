package service

import "github.com/funtimecoding/soil/pkg/linkace/tag"

func (s *Service) UpdateTag(
	identifier int,
	patch map[string]any,
) (*tag.Tag, error) {
	return guarded(
		s.guard,
		func() (*tag.Tag, error) {
			return s.client.UpdateTag(identifier, patch)
		},
	)
}
