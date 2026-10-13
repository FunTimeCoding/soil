package model_context_server

import (
	"net/http"
	"testing"
)

func (s *Server) VerifyBase(t *testing.T) {
	t.Helper()
	s.VerifyStatus(t, "/health", http.StatusOK)
	s.VerifyStatus(t, "/version", http.StatusOK)
}
