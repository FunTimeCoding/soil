package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
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
		Reporter: memory.New(),
		clock:    time.Now,
		litePath: filepath.Join(t.TempDir(), "gotelemetryd.db"),
		t:        t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
