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
	o ...gitlab.Option,
) *Tester {
	t.Helper()
	m := http.NewServeMux()
	configure(m)
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)
	result := &Tester{
		t: t,
		Client: gitlab.New(
			web.TrimScheme(s.URL),
			"test-token",
			append(
				[]gitlab.Option{gitlab.WithInsecure(), gitlab.WithRetryMax(0)},
				o...,
			)...,
		),
	}
	t.Cleanup(result.Client.Close)

	return result
}
