package base

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama/reachable_skip"
	"testing"
)

func (s *Server) SkipUnreachable(t *testing.T) {
	t.Helper()
	reachable_skip.Skip(t)
}
