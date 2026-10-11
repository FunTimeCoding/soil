package tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/url"
	"testing"
)

func (o *Tester) ExchangeCodeRaw(
	t *testing.T,
	form url.Values,
) int {
	t.Helper()
	r, e := o.client.PostForm(
		fmt.Sprintf("%s/token", o.server.BaseLocator()),
		form,
	)
	assert.FatalOnError(t, e)
	errors.PanicClose(r.Body)

	return r.StatusCode
}
