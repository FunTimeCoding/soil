package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (o *Tester) PostForm(
	path string,
	values url.Values,
) string {
	o.t.Helper()
	r, e := http.Post(
		join.Empty(o.base, path),
		constant.FormEncoded,
		strings.NewReader(values.Encode()),
	)
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(o.t, f)

	return string(b)
}
