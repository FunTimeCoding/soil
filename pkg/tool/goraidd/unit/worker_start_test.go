package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/worker"
	"path/filepath"
	"testing"
	"time"
)

func TestAWorkerStartedBeforeTheLogCacheSyncsItOnArrival(t *testing.T) {
	cache := t.TempDir()
	s := service_tester.New(t, cache, t.TempDir())
	w := worker.New(s, logger.New(context.Background()), memory.New())
	w.Start()
	defer w.Close()
	time.Sleep(50 * time.Millisecond)
	service_tester.WriteLogCache(cache, "20261024-201500.zevtc")
	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) && len(s.Fights()) == 0 {
		time.Sleep(5 * time.Millisecond)
	}

	assert.Count(t, 1, s.Fights())
}

func TestAWorkerClosesSafelyTwice(t *testing.T) {
	s := service_tester.New(t, t.TempDir(), t.TempDir())
	w := worker.New(s, logger.New(context.Background()), memory.New())
	w.Start()
	w.Close()
	w.Close()
}

func TestAClosedWorkerHasStopped(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")
	s := service_tester.New(t, missing, t.TempDir())
	r := memory.New()
	w := worker.New(s, logger.New(context.Background()), r)
	w.Start()
	w.Close()
	closed := len(r.Events())
	time.Sleep(100 * time.Millisecond)
	assert.Count(t, closed, r.Events())
}

func TestAWorkerThatCannotWatchReportsInsteadOfCrashing(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")
	s := service_tester.New(t, missing, t.TempDir())
	r := memory.New()
	w := worker.New(s, logger.New(context.Background()), r)
	w.Start()
	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) && len(r.Events()) < 2 {
		time.Sleep(5 * time.Millisecond)
	}

	w.Close()
	assert.Count(t, 2, r.Events())
}
