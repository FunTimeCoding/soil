package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func (s *Server) VerifySignIn(t *testing.T) {
	t.Helper()
	r, e := s.pool.Halting().Do(web.NewGet(s.Address(constant.SignInPath)))
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(t, http.StatusFound, r.StatusCode)

	for _, k := range r.Cookies() {
		if k.Name == constant.AuthorizationFlowCookie {
			return
		}
	}

	t.Fatalf("%s set no authorization flow cookie", constant.SignInPath)
}
