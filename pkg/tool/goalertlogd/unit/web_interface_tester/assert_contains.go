package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func (o *Tester) AssertContains(
	expected string,
	path string,
) {
	o.t.Helper()
	r, e := o.server.Client().Get(o.server.Address(path))
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(o.t, f)
	assert.StringContains(o.t, expected, string(b))
}
