package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/types/discovered_file"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/worker"
	"go.uber.org/goleak"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDaemonClosesTwice(t *testing.T) {
	d := base.New(t).Daemon()
	d.Close()
	d.Close()
}

func TestDaemonHandsOutItsWatcher(t *testing.T) {
	d := base.New(t).Daemon()
	workers := d.Workers()
	assert.Count(t, 1, workers)
	_, watcher := workers[0].(*worker.Watcher)
	assert.True(t, watcher)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := base.New(t, base.WithSeedDirectory(t.TempDir()))
	before := goleak.IgnoreCurrent()
	d := s.Daemon()
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

func TestWatcherSyncsWhatItFindsAtStart(t *testing.T) {
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

func TestWatcherClosesOnlyAfterItsSyncReturns(t *testing.T) {
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
