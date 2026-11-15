package proxmox_server

import (
	"github.com/funtimecoding/soil/pkg/proxmox/unit/proxmox_server/answer"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	result := &Server{t: t, answers: map[string]*answer.Answer{}}
	result.server = httptest.NewTLSServer(http.HandlerFunc(result.serve))
	t.Cleanup(result.server.Close)

	return result
}
