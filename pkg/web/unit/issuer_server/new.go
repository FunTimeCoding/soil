package issuer_server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	result := &Server{t: t}
	result.Server = httptest.NewServer(http.HandlerFunc(result.serve))
	t.Cleanup(result.Close)

	return result
}
