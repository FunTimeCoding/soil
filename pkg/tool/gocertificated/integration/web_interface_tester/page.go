package web_interface_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (o *Tester) Page(path string) string {
	o.t.Helper()
	r, e := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://127.0.0.1:%d%s", o.Server.Port, path),
		nil,
	)
	assert.FatalOnError(o.t, e)
	r.AddCookie(o.Server.Authorization.SubjectCookie("tester"))
	result, f := http.DefaultClient.Do(r)
	assert.FatalOnError(o.t, f)
	assert.Integer(o.t, http.StatusOK, result.StatusCode)

	return web.ReadString(result)
}
