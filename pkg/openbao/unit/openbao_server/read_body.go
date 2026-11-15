package openbao_server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"net/http"
)

func readBody(r *http.Request) []byte {
	b, e := io.ReadAll(r.Body)
	errors.PanicOnError(e)

	return b
}
