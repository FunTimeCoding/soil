package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/url"
	"testing"
)

func (s *Server) PostStatus(
	t *testing.T,
	target string,
	form url.Values,
) int {
	t.Helper()
	r, e := s.pool.Client().PostForm(s.Address(target), form)
	assert.FatalOnError(t, e)
	errors.PanicClose(r.Body)

	return r.StatusCode
}
