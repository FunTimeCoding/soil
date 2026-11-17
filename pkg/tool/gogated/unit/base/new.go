package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Server {
	t.Helper()
	result := &Server{
		Reporter: memory.New(),
		litePath: filepath.Join(t.TempDir(), "gogated.db"),
		t:        t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
