package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
	"path/filepath"
	"testing"
	"time"
)

func New(
	t *testing.T,
	o ...Option,
) *Server {
	t.Helper()
	result := &Server{
		System:   mock_system.New(t),
		Reporter: memory.New(),
		clock:    time.Now,
		litePath: filepath.Join(t.TempDir(), "goflightd.db"),
		t:        t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
