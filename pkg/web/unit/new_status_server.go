package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newStatusServer(t *testing.T) *httptest.Server {
	t.Helper()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				switch r.URL.Path {
				case "/okay":
					_, e := w.Write([]byte("alfa"))
					errors.PanicOnError(e)
				case "/missing":
					w.WriteHeader(http.StatusNotFound)
				default:
					w.WriteHeader(http.StatusInternalServerError)
				}
			},
		),
	)
	t.Cleanup(s.Close)

	return s
}
