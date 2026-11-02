package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/types/runner_option"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"sync/atomic"
	"testing"
	"time"
)

func TestLifecycleClosesNeverStartedRunner(t *testing.T) {
	var setup atomic.Bool
	l := lifecycle.New(
		logger.New(context.Background()),
	).WithWorker(newSetupProbe(&setup))
	start := time.Now()
	l.Close()
	assert.Less(t, time.Second, time.Since(start))
	assert.False(t, setup.Load())
}

func TestClosingTheRunnerEndsItsSetup(t *testing.T) {
	r := runner.New(
		runner_option.Option{
			SetupFunction: func(x context.Context) bool {
				<-x.Done()

				return false
			},
		},
		logger.New(context.Background()),
		memory.New(),
	)
	r.Start()
	r.Close()

	select {
	case <-r.Finished():
	case <-time.After(time.Second):
		t.Fatal("setup kept running after close")
	}
}

func TestATriggerAfterCloseIsRefused(t *testing.T) {
	var setup atomic.Bool
	r := newSetupProbe(&setup)
	r.Start()
	r.Close()
	awaitFinished(t, r)
	assert.True(t, unavailable.Is(r.Trigger(trigger.Request{})))
}

func TestASyncAfterCloseIsRefused(t *testing.T) {
	var setup atomic.Bool
	r := newSetupProbe(&setup)
	r.Start()
	r.Close()
	awaitFinished(t, r)
	refused := make(chan error, 1)
	go func() {
		_, e := r.Sync()
		refused <- e
	}()

	select {
	case e := <-refused:
		assert.True(t, unavailable.Is(e))
	case <-time.After(time.Second):
		t.Fatal("sync blocked on a closed runner")
	}
}

func TestATriggerQueuedBeforeStartIsAnsweredOnClose(t *testing.T) {
	var setup atomic.Bool
	r := newSetupProbe(&setup)
	response := make(chan *trigger.Result, 1)
	assert.FatalOnError(t, r.Trigger(trigger.Request{Response: response}))
	r.Close()

	select {
	case v := <-response:
		assert.True(t, unavailable.Is(v.Error))
	case <-time.After(time.Second):
		t.Fatal("a trigger queued before start was never answered")
	}
}

func TestRunnerStartAfterCloseDoesNotRun(t *testing.T) {
	var setup atomic.Bool
	r := newSetupProbe(&setup)
	r.Close()
	r.Start()
	time.Sleep(100 * time.Millisecond)
	assert.False(t, setup.Load())

	select {
	case <-r.Finished():
	default:
		t.Fatal("never-started runner did not finish on close")
	}
}
