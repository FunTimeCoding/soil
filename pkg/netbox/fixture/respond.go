package fixture

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func respond(
	w http.ResponseWriter,
	payload string,
) {
	_, e := w.Write([]byte(payload))
	errors.PanicOnError(e)
}
