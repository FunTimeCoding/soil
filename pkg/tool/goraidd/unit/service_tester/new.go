package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/client/mock_client"
	"path/filepath"
	"testing"
	"time"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	result := &Tester{
		Parser:   mock_client.New(t),
		Reporter: memory.New(),
		LogCache: t.TempDir(),
		Output:   t.TempDir(),
		Elite:    t.TempDir(),
		litePath: filepath.Join(t.TempDir(), "goraidd.db"),
		clock:    time.Now,
		t:        t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
