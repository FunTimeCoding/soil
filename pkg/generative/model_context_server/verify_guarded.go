package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
)

func (s *Server) VerifyGuarded(
	t *testing.T,
	path string,
) {
	t.Helper()
	assert.HTTPStatus(t, s.address(path), http.StatusUnauthorized)
	q := web.NewGet(s.address(path))
	web.Bearer(q, constant.ModelContextTestToken)
	r, e := web.Client().Do(q)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	assert.True(t, r.StatusCode != http.StatusUnauthorized)
}
