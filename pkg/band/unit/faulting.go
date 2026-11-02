package unit

import (
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"net/http"
	"net/http/httptest"
	"testing"
)

func faulting(
	t *testing.T,
	status int,
	body string,
) string {
	t.Helper()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				if q.Header.Get("Authorization") == "" {
					w.Header().Set(
						"WWW-Authenticate",
						constant.FixtureChallenge,
					)
					w.WriteHeader(http.StatusUnauthorized)

					return
				}

				w.Header().Set("Content-Type", constant.ContentType)
				w.WriteHeader(status)
				writer.Print(w, "%s", body)
			},
		),
	)
	t.Cleanup(s.Close)

	return s.URL
}
