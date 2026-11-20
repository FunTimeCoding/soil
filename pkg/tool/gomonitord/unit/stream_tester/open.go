package stream_tester

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (o *Tester) Open(s *model_context_server.Server) *bufio.Reader {
	o.t.Helper()
	q := web.NewGet(s.Address(constant.StreamPath))
	web.Bearer(q, generativeConstant.ModelContextTestToken)
	r, e := s.Client().Do(q)
	assert.FatalOnError(o.t, e)
	o.t.Cleanup(func() { errors.PanicClose(r.Body) })
	assert.Integer(o.t, http.StatusOK, r.StatusCode)

	return bufio.NewReader(r.Body)
}
