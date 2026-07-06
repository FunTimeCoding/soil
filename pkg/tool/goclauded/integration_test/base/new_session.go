package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/model_context_tester"
	"testing"
)

func (s *Server) NewSession(t *testing.T) *model_context_tester.Session {
	t.Helper()

	return model_context_tester.New(t, s.server.Port)
}
