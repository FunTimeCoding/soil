package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonHandsOutItsGuardsThenItsClock(t *testing.T) {
	d := daemon.New(
		service_tester.New(t).Option,
		logger.New(context.Background()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	t.Cleanup(d.Close)
	assert.Count(t, 5, d.Workers())
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	o := service_tester.New(t).Option
	before := goleak.IgnoreCurrent()
	d := daemon.New(
		o,
		logger.New(context.Background()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	d.Close()
	goleak.VerifyNone(t, before)
}
