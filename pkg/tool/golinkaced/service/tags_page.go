package service

import (
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

func (s *Service) TagsPage(p int) (*page.Page[*tag.Tag], error) {
	return guarded(
		s.guard,
		func() (*page.Page[*tag.Tag], error) { return s.client.TagsPage(p) },
	)
}
