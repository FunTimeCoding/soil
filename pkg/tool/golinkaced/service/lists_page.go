package service

import (
	"github.com/funtimecoding/soil/pkg/linkace/list"
	"github.com/funtimecoding/soil/pkg/linkace/page"
)

func (s *Service) ListsPage(p int) (*page.Page[*list.List], error) {
	return guarded(
		s.guard,
		func() (*page.Page[*list.List], error) { return s.client.ListsPage(p) },
	)
}
