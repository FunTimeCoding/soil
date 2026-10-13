package unit

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
)

func openStream(
	t *testing.T,
	s *model_context_server.Server,
) *bufio.Reader {
	t.Helper()
	q := web.NewGet(s.Address(constant.StreamPath))
	web.Bearer(q, generativeConstant.ModelContextTestToken)
	r, e := s.Client().Do(q)
	assert.FatalOnError(t, e)
	t.Cleanup(func() { errors.PanicClose(r.Body) })
	assert.Integer(t, http.StatusOK, r.StatusCode)

	return bufio.NewReader(r.Body)
}
