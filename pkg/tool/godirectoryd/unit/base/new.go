package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/service_tester"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()

	return NewWith(t, service_tester.NewSource(), memory.New())
}
