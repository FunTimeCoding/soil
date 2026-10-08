package unit

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newSaltServer(
	t *testing.T,
	password string,
) *SaltServer {
	t.Helper()
	result := &SaltServer{}
	result.Server = httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				result.serve(w, r, password)
			},
		),
	)
	t.Cleanup(result.Close)

	return result
}
