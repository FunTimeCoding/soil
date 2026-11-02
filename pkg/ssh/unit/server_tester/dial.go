package server_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/ssh"
)

func (s *Server) Dial() *ssh.Client {
	s.t.Helper()
	result := ssh.NewWithPassword("alfa", "127.0.0.1", s.Password, false).
		WithPort(s.Port)
	assert.FatalOnError(s.t, result.Dial())
	s.t.Cleanup(result.Close)

	return result
}
