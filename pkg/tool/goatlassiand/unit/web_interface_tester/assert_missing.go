package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"strings"
)

func (o *Tester) AssertMissing(
	expectedAbsent string,
	path string,
) {
	o.t.Helper()
	r, e := o.server.Client().Get(o.server.Address(path))
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(o.t, f)
	assert.True(o.t, !strings.Contains(string(b), expectedAbsent))
}
