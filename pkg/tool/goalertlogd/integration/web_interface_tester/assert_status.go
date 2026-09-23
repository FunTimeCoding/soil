package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
)

func (o *Tester) AssertStatus(
	expected int,
	path string,
) {
	o.t.Helper()
	r, e := http.Get(join.Empty(o.base, path))
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(o.t, expected, r.StatusCode)
}
