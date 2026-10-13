package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (o *Tester) Page(path string) string {
	o.t.Helper()
	r, e := http.NewRequest(http.MethodGet, o.Server.Address(path), nil)
	assert.FatalOnError(o.t, e)
	r.AddCookie(o.Server.Authorization.SubjectCookie("tester"))
	result, f := o.Server.Client().Do(r)
	assert.FatalOnError(o.t, f)
	assert.Integer(o.t, http.StatusOK, result.StatusCode)

	return web.ReadString(result)
}
