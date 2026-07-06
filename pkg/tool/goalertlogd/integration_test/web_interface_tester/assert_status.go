package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
)

func (o *Tester) AssertStatus(
	path string,
	status int,
) {
	o.t.Helper()
	r, e := http.Get(join.Empty(o.base, path))
	assert.FatalOnError(o.t, e)
	assert.Integer(o.t, status, r.StatusCode)
}
