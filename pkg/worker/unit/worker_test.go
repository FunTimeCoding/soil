package unit

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/worker"
	"sync/atomic"
	"testing"
	"time"
)

func TestOnceRunsTheTickOnce(t *testing.T) {
	var ticks atomic.Int64
	w := worker.NewOnce(
		func(context.Context) []error {
			ticks.Add(1)

			return nil
		},
		logger.New(t.Context()),
		memory.New(),
	)
	w.Start()
	w.Close()
	assert.Integer(t, int64(1), ticks.Load())
}

func TestCloseBeforeStartReturns(t *testing.T) {
	w := worker.New(
		func(context.Context) []error { return nil },
		logger.New(t.Context()),
		memory.New(),
		time.Hour,
	)
	closed := make(chan struct{})
	go func() {
		w.Close()
		w.Close()
		close(closed)
	}()

	select {
	case <-closed:
	case <-time.After(time.Second):
		t.Fatal("close before start did not return")
	}
}

func TestCloseWaitsForARunningTick(t *testing.T) {
	entered := make(chan struct{})
	var finished atomic.Bool
	w := worker.NewOnce(
		func(x context.Context) []error {
			close(entered)
			<-x.Done()
			finished.Store(true)

			return nil
		},
		logger.New(t.Context()),
		memory.New(),
	)
	w.Start()
	<-entered
	w.Close()
	assert.True(t, finished.Load())
}

func TestAReportableFailureIsCapturedAndAQuietOneIsNot(t *testing.T) {
	r := memory.New()
	worker.NewEvery(
		func() []error {
			return []error{fmt.Errorf("broken"), validation.New("refused")}
		},
		logger.New(t.Context()),
		r,
		time.Hour,
	).Poll()
	assert.Count(t, 1, r.Events())
}
