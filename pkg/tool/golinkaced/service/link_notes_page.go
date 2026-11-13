package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/note"
	"github.com/funtimecoding/soil/pkg/linkace/page"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) LinkNotesPage(
	linkIdentifier int,
	p int,
) (*page.Page[*note.Note], error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*page.Page[*note.Note], error) {
			return s.client.NotesByLinkPage(linkIdentifier, p)
		},
	)
}
