package upstream_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func answer(
	w http.ResponseWriter,
	status int,
	body string,
) {
	w.WriteHeader(status)
	_, e := w.Write([]byte(body))
	errors.PanicOnError(e)
}
