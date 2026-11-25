package note_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
)

func (s *Server) Notes(after int64) []*note.Note {
	s.t.Helper()
	result, e := s.client.MergeRequestNotes(s.t.Context(), 7, 3, after)
	assert.FatalOnError(s.t, e)

	return result
}
