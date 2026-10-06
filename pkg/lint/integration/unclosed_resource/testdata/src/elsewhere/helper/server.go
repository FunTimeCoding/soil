package helper

import (
	"net/http"
	"testing"
)

type Server struct {
	response *http.Response
}

func (s *Server) Close() {
	_ = s.response.Body.Close()
}

func (s *Server) Ping() {}

func NewServer(t *testing.T) *Server {
	r, _ := http.Get("http://example.com")
	result := &Server{response: r}
	t.Cleanup(result.Close)

	return result
}
