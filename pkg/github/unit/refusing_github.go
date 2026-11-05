package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github"
	"net/http"
	"net/http/httptest"
	"testing"
)

func refusingGithub(
	t *testing.T,
	status int,
	header http.Header,
	body string,
) *github.Client {
	t.Helper()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				for key, values := range header {
					w.Header()[key] = values
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(status)
				_, e := w.Write([]byte(body))
				errors.PanicOnError(e)
			},
		),
	)
	t.Cleanup(s.Close)
	result := github.NewServer("token", s.URL)
	t.Cleanup(result.Close)

	return result
}
