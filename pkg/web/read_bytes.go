package web

import (
	"github.com/funtimecoding/soil/pkg/errors/unreadable_body"
	"io"
	"net/http"
)

func ReadBytes(r *http.Response) []byte {
	result, e := io.ReadAll(r.Body)

	if e != nil {
		panic(unreadable_body.New(e, "read response body"))
	}

	return result
}
