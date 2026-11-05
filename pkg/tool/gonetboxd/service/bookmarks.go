package service

import "github.com/funtimecoding/soil/pkg/netbox/bookmark"

func (s *Service) Bookmarks() ([]*bookmark.Bookmark, error) {
	return guarded(
		s.guard,
		func() ([]*bookmark.Bookmark, error) {
			return s.client.Bookmarks()
		},
	)
}
