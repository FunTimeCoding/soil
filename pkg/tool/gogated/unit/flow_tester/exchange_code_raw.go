package flow_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/url"
)

func (o *Tester) ExchangeCodeRaw(form url.Values) int {
	o.t.Helper()
	r, e := o.client.PostForm(
		fmt.Sprintf("%s/token", o.server.BaseLocator()),
		form,
	)
	assert.FatalOnError(o.t, e)
	errors.PanicClose(r.Body)

	return r.StatusCode
}
