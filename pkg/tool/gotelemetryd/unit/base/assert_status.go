package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Server) AssertStatus(
	expected int,
	path string,
) {
	s.t.Helper()
	v := s.Served()
	r, e := v.Client().Get(v.Address(path))
	assert.FatalOnError(s.t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(s.t, expected, r.StatusCode)
}
