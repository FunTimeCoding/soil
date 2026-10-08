package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
)

func (s *Server) VerifySession(
	t *testing.T,
	path string,
) {
	t.Helper()
	bare, e := web.HaltingClient().Do(web.NewGet(s.address(path)))
	assert.FatalOnError(t, e)
	errors.PanicClose(bare.Body)
	assert.Integer(t, http.StatusFound, bare.StatusCode)
	q := web.NewGet(s.address(path))
	web.Bearer(q, constant.ModelContextTestToken)
	r, f := web.HaltingClient().Do(q)
	assert.FatalOnError(t, f)
	defer errors.PanicClose(r.Body)
	assert.Integer(t, http.StatusOK, r.StatusCode)
}
