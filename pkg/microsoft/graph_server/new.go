package graph_server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	result := &Server{answers: map[string][]Answer{}}
	result.server = httptest.NewServer(http.HandlerFunc(result.serve))
	t.Cleanup(result.server.Close)

	return result
}
