package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/types/discovered_file"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/worker"
	"go.uber.org/goleak"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestTheDaemonClosesTwice(t *testing.T) {
	d := daemon.New(
		store_tester.NewOption(t, time.Now, notifier.New()),
		logger.New(t.Context()),
		memory.New(),
	)
	d.Close()
	d.Close()
}

func TestTheDaemonHandsOutItsWatcher(t *testing.T) {
	d := daemon.New(
		store_tester.NewOption(t, time.Now, notifier.New()),
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(d.Close)
	workers := d.Workers()
	assert.Count(t, 1, workers)
	_, watcher := workers[0].(*worker.Watcher)
	assert.True(t, watcher)
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	o := store_tester.NewOption(t, time.Now, notifier.New())
	o.SeedDirectory = t.TempDir()
	d := daemon.New(o, logger.New(t.Context()), memory.New())
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}

func TestTheWatcherSyncsWhatItFindsAtStart(t *testing.T) {
	directory := t.TempDir()
	assert.FatalOnError(
		t,
		os.WriteFile(
			filepath.Join(directory, "alfa.md"),
			[]byte("# Alfa"),
			0o644,
		),
	)
	found := make(chan []discovered_file.File, 1)
	w := worker.New(
		func(files []discovered_file.File) { found <- files },
		directory,
		logger.New(t.Context()),
		memory.New(),
	)
	w.Start()
	defer w.Close()

	select {
	case files := <-found:
		assert.Count(t, 1, files)
		assert.String(t, "alfa", files[0].Name)
	case <-time.After(5 * time.Second):
		t.Fatal("the watcher synced nothing")
	}
}

func TestTheWatcherClosesOnlyAfterItsSyncReturns(t *testing.T) {
	entered := make(chan struct{})
	release := make(chan struct{})
	w := worker.New(
		func([]discovered_file.File) {
			close(entered)
			<-release
		},
		t.TempDir(),
		logger.New(t.Context()),
		memory.New(),
	)
	w.Start()
	<-entered
	closed := make(chan struct{})
	go func() {
		w.Close()
		close(closed)
	}()

	select {
	case <-closed:
		t.Fatal("the watcher closed while its sync was in flight")
	case <-time.After(100 * time.Millisecond):
	}

	close(release)
	<-closed
}
