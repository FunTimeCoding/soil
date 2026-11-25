package note_server

import "github.com/funtimecoding/soil/pkg/assert"

func (s *Server) Approvals() []string {
	s.t.Helper()
	result, e := s.client.MergeRequestApprovals(s.t.Context(), 7, 3)
	assert.FatalOnError(s.t, e)

	return result
}
