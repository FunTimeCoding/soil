package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (o *Tester) AssertStatus(
	expected int,
	path string,
) {
	o.t.Helper()
	r, e := o.server.Client().Get(o.server.Address(path))
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(o.t, expected, r.StatusCode)
}
