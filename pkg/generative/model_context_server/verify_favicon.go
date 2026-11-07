package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func (s *Server) VerifyFavicon(t *testing.T) {
	t.Helper()
	r, e := s.pool.Client().Do(web.NewGet(s.Address(webConstant.FaviconPath)))
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(t, http.StatusOK, r.StatusCode)
	assert.String(t, "image/png", r.Header.Get(webConstant.ContentType))
}
