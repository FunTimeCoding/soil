package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestACancelledCleanupRemovesNothing(t *testing.T) {
	o := service_tester.NewOption(
		t,
		t.TempDir(),
		t.TempDir(),
		service_tester.NewParser(t),
	)
	old := filepath.Join(o.ElitePath, "old.json")
	assert.FatalOnError(t, os.WriteFile(old, nil, 0o600))
	month := time.Now().AddDate(0, -1, 0)
	assert.FatalOnError(t, os.Chtimes(old, month, month))
	s := service_tester.NewService(t, o, memory.New())
	x, cancel := context.WithCancel(t.Context())
	cancel()
	assert.Count(t, 0, s.Cleanup(x))
	_, e := os.Stat(old)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, s.Cleanup(t.Context()))
	_, f := os.Stat(old)
	assert.True(t, os.IsNotExist(f))
}
