package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/openbao/status"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/watch"
	"github.com/prometheus/client_golang/prometheus"
	"sync/atomic"
	"testing"
	"time"
)

func TestWatchReportsAnUnsealedStore(t *testing.T) {
	y := prometheus.NewRegistry()
	watch.New(
		service_tester.New(t, mock_client.New()).Status,
		y,
		logger.New(context.Background()),
		memory.New(),
	).Poll()
	assert.Float(t, 0, gauge(t, y, constant.SealedGauge))
	assert.Float(t, 1, gauge(t, y, constant.ReachableGauge))
}

func TestWatchReportsASealedStore(t *testing.T) {
	y := prometheus.NewRegistry()
	c := mock_client.New()
	c.Seal()
	watch.New(
		service_tester.New(t, c).Status,
		y,
		logger.New(context.Background()),
		memory.New(),
	).Poll()
	assert.Float(t, 1, gauge(t, y, constant.SealedGauge))
}

func TestWatchReportsAnUnreachableStore(t *testing.T) {
	y := prometheus.NewRegistry()
	c := mock_client.New()
	c.FailHealth(errors.New("connection refused"))
	watch.New(
		service_tester.New(t, c).Status,
		y,
		logger.New(context.Background()),
		memory.New(),
	).Poll()
	assert.Float(t, 0, gauge(t, y, constant.ReachableGauge))
}

func TestWatchCloseWaitsForATickInFlight(t *testing.T) {
	entered := make(chan struct{})
	release := make(chan struct{})
	var finished atomic.Bool
	w := watch.New(
		func() (*status.Status, error) {
			close(entered)
			<-release
			finished.Store(true)

			return status.New(true, false, ""), nil
		},
		prometheus.NewRegistry(),
		logger.New(context.Background()),
		memory.New(),
	)
	w.Start()
	<-entered
	closed := make(chan bool)
	go func() {
		w.Close()
		closed <- finished.Load()
	}()
	time.Sleep(50 * time.Millisecond)
	close(release)
	assert.True(t, <-closed)
}
