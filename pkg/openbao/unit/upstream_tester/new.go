package upstream_tester

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	result := &Tester{}
	result.server = httptest.NewServer(http.HandlerFunc(result.serve))
	t.Cleanup(result.server.Close)

	return result
}
