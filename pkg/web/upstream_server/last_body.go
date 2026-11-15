package upstream_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"io"
)

func (s *Server) LastBody() string {
	s.t.Helper()
	b, e := io.ReadAll(s.Last().Body)
	assert.FatalOnError(s.t, e)

	return string(b)
}
