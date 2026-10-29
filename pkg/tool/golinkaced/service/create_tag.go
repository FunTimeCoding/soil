package service

import "github.com/funtimecoding/soil/pkg/linkace/tag"

func (s *Service) CreateTag(name string) (*tag.Tag, error) {
	return guarded(
		s.guard,
		func() (*tag.Tag, error) { return s.client.CreateTag(name) },
	)
}
