package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"io"
	"net/http"
)

func (o *Tester) AssertContains(
	path string,
	expected string,
) {
	o.t.Helper()
	r, e := http.Get(join.Empty(o.base, path))
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(o.t, f)
	assert.StringContains(o.t, expected, string(b))
}
