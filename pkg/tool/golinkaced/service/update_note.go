package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/note"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) UpdateNote(
	identifier int,
	text string,
) (*note.Note, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*note.Note, error) {
			return s.client.UpdateNote(identifier, text)
		},
	)
}
