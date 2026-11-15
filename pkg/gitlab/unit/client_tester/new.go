package client_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	configure func(*http.ServeMux),
	o ...gitlab.Option,
) *Tester {
	t.Helper()
	m := http.NewServeMux()
	configure(m)
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)

	return &Tester{t: t, Client: Connect(t, s.URL, o...)}
}
