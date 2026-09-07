package base

import "testing"

func (s *Server) SkipUnreachable(t *testing.T) {
	t.Helper()

	if !s.embedder.Reachable() {
		t.Skipf("embed host unreachable: %s", s.embedder.Locator())
	}
}
