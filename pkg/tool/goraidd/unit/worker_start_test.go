package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestAWorkerStartedBeforeTheLogCacheSyncsItOnArrival(t *testing.T) {
	cache := t.TempDir()
	s := service_tester.New(t, cache, t.TempDir())
	w := service_tester.NewWatcher(t, s, memory.New())
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
	w := service_tester.NewWatcher(t, s, memory.New())
	w.Start()
	w.Close()
	w.Close()
}

func TestAClosedWorkerHasStopped(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")
	s := service_tester.New(t, missing, t.TempDir())
	r := memory.New()
	w := service_tester.NewWatcher(t, s, r)
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
	w := service_tester.NewWatcher(t, s, r)
	w.Start()
	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) && len(r.Events()) < 2 {
		time.Sleep(5 * time.Millisecond)
	}

	w.Close()
	assert.Count(t, 2, r.Events())
}

func TestAFailedStartUpStepLeavesTheNextOnesRunning(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")
	o := service_tester.NewOption(
		t,
		missing,
		t.TempDir(),
		service_tester.NewParser(t),
	)
	expired := filepath.Join(
		o.ElitePath,
		"20260101-120000_detailed_wvw_kill.json",
	)
	assert.FatalOnError(t, os.WriteFile(expired, []byte("{}"), 0o600))
	old := time.Now().AddDate(0, 0, -30)
	assert.FatalOnError(t, os.Chtimes(expired, old, old))
	s := service_tester.NewService(t, o, memory.New())
	w := service_tester.NewWatcher(t, s, memory.New())
	w.Start()
	defer w.Close()
	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) {
		if _, e := os.Stat(expired); os.IsNotExist(e) {
			return
		}

		time.Sleep(5 * time.Millisecond)
	}

	t.Fatal("the cleanup after a failed log sync never ran")
}
