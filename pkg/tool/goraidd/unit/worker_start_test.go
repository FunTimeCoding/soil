package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWorkerStartedBeforeLogCacheSyncsItOnArrival(t *testing.T) {
	o := service_tester.New(t)
	s := o.Service()
	w := o.Watcher(s)
	w.Start()
	defer w.Close()
	time.Sleep(50 * time.Millisecond)
	service_tester.WriteLogCache(o.LogCache, "20261024-201500.zevtc")
	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) && len(s.Fights()) == 0 {
		time.Sleep(5 * time.Millisecond)
	}

	assert.Count(t, 1, s.Fights())
}

func TestWorkerClosesSafelyTwice(t *testing.T) {
	o := service_tester.New(t)
	w := o.Watcher(o.Service())
	w.Start()
	w.Close()
	w.Close()
}

func TestClosedWorkerHasStopped(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithLogCache(filepath.Join(t.TempDir(), "missing")),
	)
	w := o.Watcher(o.Service())
	w.Start()
	w.Close()
	closed := len(o.Reporter.Events())
	time.Sleep(100 * time.Millisecond)
	assert.Count(t, closed, o.Reporter.Events())
}

func TestWorkerThatCannotWatchReportsInsteadOfCrashing(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithLogCache(filepath.Join(t.TempDir(), "missing")),
	)
	w := o.Watcher(o.Service())
	w.Start()
	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) && len(o.Reporter.Events()) < 2 {
		time.Sleep(5 * time.Millisecond)
	}

	w.Close()
	assert.Count(t, 2, o.Reporter.Events())
}

func TestFailedStartUpStepLeavesNextOnesRunning(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithLogCache(filepath.Join(t.TempDir(), "missing")),
	)
	expired := filepath.Join(o.Elite, "20260101-120000_detailed_wvw_kill.json")
	assert.FatalOnError(t, os.WriteFile(expired, []byte("{}"), 0o600))
	old := time.Now().AddDate(0, 0, -30)
	assert.FatalOnError(t, os.Chtimes(expired, old, old))
	w := o.Watcher(o.Service())
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
