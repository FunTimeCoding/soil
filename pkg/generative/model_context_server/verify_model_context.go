package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"net/http"
	"testing"
)

func (s *Server) VerifyModelContext(t *testing.T) {
	t.Helper()
	s.VerifyStatus(t, "/mcp", http.StatusUnauthorized)
	s.VerifyStatus(t, "/sse", http.StatusUnauthorized)
	c := s.Connect(t)
	defer c.Close()
	assert.True(t, len(c.ListTools()) > 0)
}
