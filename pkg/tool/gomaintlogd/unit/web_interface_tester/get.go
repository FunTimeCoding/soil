package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func (o *Tester) Get(path string) string {
	o.t.Helper()
	r, e := o.server.Client().Get(o.server.Address(path))
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(o.t, f)

	return string(b)
}
