package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
)

func (s *Server) VerifyOpenPost(
	t *testing.T,
	path string,
) {
	t.Helper()
	q := web.NewPost(s.address(path), "{}")
	r, e := web.Client().Do(q)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	assert.True(t, r.StatusCode != http.StatusUnauthorized)
}
