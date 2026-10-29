package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func (s *Server) VerifyBase(t *testing.T) {
	t.Helper()
	s.VerifyStatus(t, webConstant.HealthPath, http.StatusOK)
	s.VerifyStatus(t, webConstant.VersionPath, http.StatusOK)
	s.VerifyStatus(t, webConstant.ReachablePath, http.StatusUnauthorized)
	q := web.NewGet(s.Address(webConstant.ReachablePath))
	web.Bearer(q, constant.ModelContextTestToken)
	r, e := s.pool.Client().Do(q)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(t, http.StatusNoContent, r.StatusCode)
}
