package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
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
		Indexer:  mock_client.New(t),
		Reporter: memory.New(),
		Path:     filepath.Join(t.TempDir(), "memory.db"),
		clock:    time.Now,
		t:        t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
