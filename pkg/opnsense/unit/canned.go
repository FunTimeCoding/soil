package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
	"net/http/httptest"
)

func canned(body string) *httptest.Server {
	return httptest.NewTLSServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				_, e := w.Write([]byte(body))
				errors.PanicOnError(e)
			},
		),
	)
}
