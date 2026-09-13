package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"io"
	"net/http"
	"strings"
)

func (o *Tester) AssertMissing(
	path string,
	absent string,
) {
	o.t.Helper()
	r, e := http.Get(join.Empty(o.base, path))
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(o.t, f)
	assert.True(o.t, !strings.Contains(string(b), absent))
}
