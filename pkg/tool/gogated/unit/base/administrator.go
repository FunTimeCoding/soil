package base

import "github.com/funtimecoding/soil/pkg/assert"

func (s *Server) Administrator(subject string) bool {
	s.t.Helper()
	result, e := s.Service().IsAdministrator(subject)
	assert.FatalOnError(s.t, e)

	return result
}
