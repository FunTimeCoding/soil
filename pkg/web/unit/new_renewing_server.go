package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newRenewingServer(
	t *testing.T,
	accepted string,
) *httptest.Server {
	t.Helper()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				if r.Header.Get(constant.Authorization) != accepted {
					w.WriteHeader(http.StatusUnauthorized)

					return
				}

				_, e := w.Write([]byte(`{"name":"alfa"}`))
				errors.PanicOnError(e)
			},
		),
	)
	t.Cleanup(s.Close)

	return s
}
