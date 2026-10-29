package service

import "github.com/funtimecoding/soil/pkg/linkace/note"

func (s *Service) UpdateNote(
	identifier int,
	text string,
) (*note.Note, error) {
	return guarded(
		s.guard,
		func() (*note.Note, error) {
			return s.client.UpdateNote(identifier, text)
		},
	)
}
