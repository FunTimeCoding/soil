package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"net/url"
	"testing"
)

func (s *Server) PostBody(
	t *testing.T,
	target string,
	form url.Values,
) string {
	t.Helper()
	r, e := s.pool.Client().PostForm(s.Address(target), form)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)

	return string(b)
}
