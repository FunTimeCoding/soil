package model_context_server

import "testing"

func (s *Server) VerifyInterface(t *testing.T) {
	t.Helper()
	s.VerifyGuarded(t, "/api/guard-probe")
}
