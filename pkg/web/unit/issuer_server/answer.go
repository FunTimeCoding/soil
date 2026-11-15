package issuer_server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func answer(
	w http.ResponseWriter,
	s string,
) {
	_, e := w.Write([]byte(s))
	errors.PanicOnError(e)
}
