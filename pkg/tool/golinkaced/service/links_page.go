package service

import (
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/funtimecoding/soil/pkg/linkace/page"
)

func (s *Service) LinksPage(p int) (*page.Page[*link.Link], error) {
	return guarded(
		s.guard,
		func() (*page.Page[*link.Link], error) { return s.client.LinksPage(p) },
	)
}
