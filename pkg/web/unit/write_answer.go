package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func writeAnswer(
	w http.ResponseWriter,
	s string,
) {
	_, e := w.Write([]byte(s))
	errors.PanicOnError(e)
}
