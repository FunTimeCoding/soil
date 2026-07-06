package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func WriteBytes(
	w http.ResponseWriter,
	code int,
	b []byte,
) int {
	w.WriteHeader(code)
	result, e := w.Write(b)
	errors.PanicOnError(e)

	return result
}
