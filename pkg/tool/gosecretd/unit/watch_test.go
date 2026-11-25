package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric/metric_tester"
	"github.com/funtimecoding/soil/pkg/openbao/status"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/watch"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"sync/atomic"
	"testing"
	"time"
)

func TestWatchReportsUnsealedStore(t *testing.T) {
	y := prometheus.NewRegistry()
	owned := watch.New(
		service_tester.New(t).Service().Status,
		y,
		logger.New(context.Background()),
		memory.New(),
	)
	defer owned.Close()
	owned.Poll()
	assert.Float(
		t,
		0,
		metric_tester.New(t).Sample(y, constant.SealedGauge, nil),
	)
	assert.Float(
		t,
		1,
		metric_tester.New(t).Sample(y, constant.ReachableGauge, nil),
	)
}

func TestWatchReportsSealedStore(t *testing.T) {
	y := prometheus.NewRegistry()
	o := service_tester.New(t)
	o.Source.Seal()
	owned := watch.New(
		o.Service().Status,
		y,
		logger.New(context.Background()),
		memory.New(),
	)
	defer owned.Close()
	owned.Poll()
	assert.Float(
		t,
		1,
		metric_tester.New(t).Sample(y, constant.SealedGauge, nil),
	)
}

func TestWatchReportsUnreachableStore(t *testing.T) {
	y := prometheus.NewRegistry()
	o := service_tester.New(t)
	o.Source.FailHealth(errors.New("connection refused"))
	owned := watch.New(
		o.Service().Status,
		y,
		logger.New(context.Background()),
		memory.New(),
	)
	defer owned.Close()
	owned.Poll()
	assert.Float(
		t,
		0,
		metric_tester.New(t).Sample(y, constant.ReachableGauge, nil),
	)
}

func TestWatchReportsSealedStoreWhileGuardIsOpen(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	answer, e := s.GeneratedClient().ListSecretsWithResponse(
		t.Context(),
		&client.ListSecretsParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode())
	s.Source.Failure = nil
	s.Source.Seal()
	s.Watch().Poll()
	assert.Float(
		t,
		1,
		metric_tester.New(t).Sample(s.Registry, constant.SealedGauge, nil),
	)
	assert.Float(
		t,
		1,
		metric_tester.New(t).Sample(s.Registry, constant.ReachableGauge, nil),
	)
}

func TestWatchCloseWaitsForTickInFlight(t *testing.T) {
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
