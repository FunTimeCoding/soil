package gopostgres

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/terminal"
	"io"
	"net/http"
)

func printResponse(
	t *terminal.Terminal,
	r *http.Response,
) {
	defer errors.PanicClose(r.Body)
	b, e := io.ReadAll(r.Body)
	errors.PanicOnError(e)

	if r.StatusCode >= http.StatusBadRequest {
		t.Reject(r.Status, b)

		return
	}

	var v any
	errors.PanicOnError(json.Unmarshal(b, &v))
	console.Line(notation.MarshalIndent(v))
}
