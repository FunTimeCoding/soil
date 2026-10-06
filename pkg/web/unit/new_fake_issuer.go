package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newFakeIssuer(t *testing.T) *fakeIssuer {
	t.Helper()
	result := &fakeIssuer{}
	result.Server = httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				if r.URL.Path == "/token" {
					errors.PanicOnError(r.ParseForm())

					if r.Form.Get("code") == "expired" {
						w.WriteHeader(http.StatusBadRequest)
						writeAnswer(
							w,
							`{"error":"invalid_grant","error_description":"Code expired"}`,
						)

						return
					}

					writeAnswer(w, `{"id_token":"alfa.bravo.charlie"}`)

					return
				}

				result.mutex.Lock()
				result.discovery++
				result.mutex.Unlock()
				w.WriteHeader(http.StatusServiceUnavailable)
			},
		),
	)
	t.Cleanup(result.Close)

	return result
}
