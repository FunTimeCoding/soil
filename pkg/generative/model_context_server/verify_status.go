package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"testing"
)

func (s *Server) VerifyStatus(
	t *testing.T,
	path string,
	status int,
) {
	t.Helper()
	r, e := s.pool.Client().Do(web.NewGet(s.Address(path)))
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(t, status, r.StatusCode)
}
