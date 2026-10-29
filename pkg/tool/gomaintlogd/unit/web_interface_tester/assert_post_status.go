package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/url"
	"strings"
)

func (o *Tester) AssertPostStatus(
	expected int,
	path string,
	values url.Values,
) {
	o.t.Helper()
	r, e := o.server.Client().Post(
		o.server.Address(path),
		constant.FormEncoded,
		strings.NewReader(values.Encode()),
	)
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(o.t, expected, r.StatusCode)
}
