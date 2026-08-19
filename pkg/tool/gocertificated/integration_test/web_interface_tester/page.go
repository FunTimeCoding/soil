package web_interface_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
)

func (o *Tester) Page(
	t *testing.T,
	path string,
) string {
	t.Helper()
	result, e := http.Get(
		fmt.Sprintf("http://127.0.0.1:%d%s", o.Server.Port(), path),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, result.StatusCode)

	return web.ReadString(result)
}
