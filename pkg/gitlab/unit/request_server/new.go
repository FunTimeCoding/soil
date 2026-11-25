package request_server

import (
	"github.com/funtimecoding/soil/pkg/gitlab/unit/client_tester"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	total int,
) *Server {
	t.Helper()
	result := &Server{t: t, total: total}
	result.client = client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc("/api/v4/projects/7/merge_requests", result.serve)
		},
	).Client

	return result
}
