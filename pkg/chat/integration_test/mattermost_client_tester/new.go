package mattermost_client_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	configure func(*http.ServeMux),
	options ...mattermost.Option,
) *Tester {
	t.Helper()
	m := http.NewServeMux()
	m.HandleFunc("/api/v4/websocket", socket)
	configure(m)
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)

	return &Tester{
		Client: mattermost.New(
			append(
				[]mattermost.Option{
					mattermost.WithHost(web.TrimScheme(s.URL)),
					mattermost.WithToken("test-token"),
					mattermost.WithInsecure(),
				},
				options...,
			)...,
		),
	}
}
