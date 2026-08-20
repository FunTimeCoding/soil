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
	r, e := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://127.0.0.1:%d%s", o.Server.Port(), path),
		nil,
	)
	assert.FatalOnError(t, e)
	r.AddCookie(o.Server.Authorization.SubjectCookie("tester"))
	result, f := http.DefaultClient.Do(r)
	assert.FatalOnError(t, f)
	assert.Integer(t, http.StatusOK, result.StatusCode)

	return web.ReadString(result)
}
