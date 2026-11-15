package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"io"
	"testing"
)

func (s *Server) Body(
	t *testing.T,
	target string,
) string {
	t.Helper()
	r, e := s.pool.Client().Do(web.NewGet(s.Address(target)))
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)

	return string(b)
}
