package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unreadable_body"
	"io"
	"net/http"
)

func ReadBytes(r *http.Response) []byte {
	defer errors.PanicClose(r.Body)
	result, e := io.ReadAll(r.Body)

	if e != nil {
		panic(unreadable_body.New(e, "read response body"))
	}

	return result
}
