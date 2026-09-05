package client_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	configure func(*http.ServeMux),
) *Tester {
	t.Helper()
	m := http.NewServeMux()
	m.HandleFunc("/api/v4/user", user)
	configure(m)
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)

	return &Tester{
		Client: gitlab.New(
			web.TrimScheme(s.URL),
			"test-token",
			gitlab.WithInsecure(),
		),
	}
}
