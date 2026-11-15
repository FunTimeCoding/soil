package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"testing"
)

func (s *Server) Status(
	t *testing.T,
	target string,
) int {
	t.Helper()
	r, e := s.pool.Client().Do(web.NewGet(s.Address(target)))
	assert.FatalOnError(t, e)
	errors.PanicClose(r.Body)

	return r.StatusCode
}
