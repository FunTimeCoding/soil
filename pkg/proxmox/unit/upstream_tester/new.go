package upstream_tester

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(t *testing.T) *Upstream {
	t.Helper()
	result := &Upstream{t: t, answers: map[string]*Answer{}}
	result.server = httptest.NewTLSServer(http.HandlerFunc(result.serve))
	t.Cleanup(result.server.Close)

	return result
}
