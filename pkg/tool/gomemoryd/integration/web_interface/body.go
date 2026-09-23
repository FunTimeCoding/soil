package web_interface

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
	"testing"
)

func body(
	t *testing.T,
	address string,
	extended bool,
) string {
	t.Helper()
	q, e := http.NewRequest(http.MethodGet, address, nil)
	assert.FatalOnError(t, e)

	if extended {
		q.Header.Set(constant.ExtendedRequest, "true")
	}

	r, f := http.DefaultClient.Do(q)
	assert.FatalOnError(t, f)
	defer errors.PanicClose(r.Body)
	b, g := io.ReadAll(r.Body)
	assert.FatalOnError(t, g)

	return string(b)
}
