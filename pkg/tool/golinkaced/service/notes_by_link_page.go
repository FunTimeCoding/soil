package service

import (
	"github.com/funtimecoding/soil/pkg/linkace/note"
	"github.com/funtimecoding/soil/pkg/linkace/page"
)

func (s *Service) NotesByLinkPage(
	linkIdentifier int,
	p int,
) (*page.Page[*note.Note], error) {
	return guarded(
		s.guard,
		func() (*page.Page[*note.Note], error) {
			return s.client.NotesByLinkPage(linkIdentifier, p)
		},
	)
}
