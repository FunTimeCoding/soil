package salt_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func write(
	w http.ResponseWriter,
	s string,
) {
	_, e := w.Write([]byte(s))
	errors.PanicOnError(e)
}
