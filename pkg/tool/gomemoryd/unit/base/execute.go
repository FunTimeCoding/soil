package base

import "github.com/funtimecoding/soil/pkg/assert"

func (s *Server) Execute(statement string) {
	_, e := s.database.Exec(statement)
	assert.FatalOnError(s.t, e)
}
