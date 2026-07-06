package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func WriteBytesSafe(
	w http.ResponseWriter,
	code int,
	b []byte,
) int {
	w.WriteHeader(code)
	result, e := w.Write(b)
	errors.LogOnError(e)

	return result
}
