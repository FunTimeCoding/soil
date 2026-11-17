package base

import "github.com/funtimecoding/soil/pkg/assert"

func (s *Server) Allows(identifier string) bool {
	s.t.Helper()
	result, e := s.Service().DirectoryAllows(identifier)
	assert.FatalOnError(s.t, e)

	return result
}
