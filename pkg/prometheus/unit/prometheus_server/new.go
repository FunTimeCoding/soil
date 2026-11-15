package prometheus_server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	status int,
	body string,
) *Server {
	t.Helper()
	result := &Server{status: status, body: body}
	result.Server = httptest.NewServer(http.HandlerFunc(result.serve))
	t.Cleanup(result.Close)

	return result
}
