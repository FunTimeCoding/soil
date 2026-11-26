package discussion_server

import (
	"github.com/funtimecoding/soil/pkg/gitlab/unit/client_tester"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	result := &Server{t: t}
	result.client = client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc("/", result.serve)
		},
	).Client

	return result
}
