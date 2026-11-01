package fake_assistant

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	result := &Server{
		dialed:     make(chan struct{}, 64),
		subscribed: make(chan struct{}, 64),
	}
	result.server = httptest.NewServer(http.HandlerFunc(result.accept))
	t.Cleanup(result.server.Close)

	return result
}
