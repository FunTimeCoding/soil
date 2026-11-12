package upstream_tester

import (
	"github.com/funtimecoding/soil/pkg/proxmox/unit/upstream_tester/answer"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(t *testing.T) *Upstream {
	t.Helper()
	result := &Upstream{t: t, answers: map[string]*answer.Answer{}}
	result.server = httptest.NewTLSServer(http.HandlerFunc(result.serve))
	t.Cleanup(result.server.Close)

	return result
}
