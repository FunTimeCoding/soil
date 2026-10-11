package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"net/http"
	"testing"
)

func body(
	t *testing.T,
	address string,
) string {
	t.Helper()
	r, e := http.Get(address)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)

	return string(b)
}
