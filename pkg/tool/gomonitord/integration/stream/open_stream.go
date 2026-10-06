package stream

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	monitorConstant "github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
)

func openStream(
	t *testing.T,
	port int,
) *bufio.Reader {
	t.Helper()
	q := web.NewGet(
		fmt.Sprintf("http://localhost:%d%s", port, monitorConstant.StreamPath),
	)
	web.Bearer(q, constant.ModelContextTestToken)
	r, e := web.Client().Do(q)
	assert.FatalOnError(t, e)
	t.Cleanup(func() { errors.PanicClose(r.Body) })
	assert.Integer(t, http.StatusOK, r.StatusCode)

	return bufio.NewReader(r.Body)
}
