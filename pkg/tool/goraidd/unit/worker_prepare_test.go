package unit

import (
	"context"
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/worker"
	"sync/atomic"
	"testing"
	"time"
)

func TestCloseCancelsStartUpStepInFlightAndSkipsRest(t *testing.T) {
	entered := make(chan struct{})
	seen := make(chan error, 1)
	var later atomic.Int64
	w := worker.New(
		func(context.Context, fsnotify.Event) []error { return nil },
		nil,
		logger.New(t.Context()),
		memory.New(),
		func(x context.Context) []error {
			close(entered)

			select {
			case <-x.Done():
				seen <- x.Err()
			case <-time.After(time.Second):
				seen <- nil
			}

			return nil
		},
		func(context.Context) []error {
			later.Add(1)

			return nil
		},
	)
	w.Start()
	<-entered
	w.Close()
	assert.True(t, errors.Is(<-seen, context.Canceled))
	assert.Integer(t, int64(0), later.Load())
}
