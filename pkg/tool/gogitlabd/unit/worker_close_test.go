package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
	"testing"
	"time"
)

func TestClosingTheClockWaitsForATickInFlight(t *testing.T) {
	entered := make(chan struct{})
	release := make(chan struct{})
	var once sync.Once
	w := worker.New(
		func() []error {
			once.Do(func() { close(entered) })
			<-release

			return nil
		},
		func() []*board_entry.Entry { return nil },
		time.Hour,
		prometheus.NewRegistry(),
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
		t.Fatal("Close returned while a tick was in flight")
	case <-time.After(100 * time.Millisecond):
	}

	close(release)

	select {
	case <-closed:
	case <-time.After(time.Second):
		t.Fatal("Close did not return after the tick ended")
	}
}

func TestClosingAClockThatNeverStartedReturns(t *testing.T) {
	w := worker.New(
		func() []error { return nil },
		func() []*board_entry.Entry { return nil },
		time.Hour,
		prometheus.NewRegistry(),
		logger.New(t.Context()),
		memory.New(),
	)
	w.Close()
	w.Close()
}
