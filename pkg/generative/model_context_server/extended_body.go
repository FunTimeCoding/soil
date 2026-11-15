package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"testing"
)

func (s *Server) ExtendedBody(
	t *testing.T,
	target string,
) string {
	t.Helper()
	q := web.NewGet(s.Address(target))
	q.Header.Set(webConstant.ExtendedRequest, "true")
	r, e := s.pool.Client().Do(q)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)

	return string(b)
}
