package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/service_tester"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	o := service_tester.NewOption(t)
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

func TestDaemonCountsReapsOnItsRegistry(t *testing.T) {
	y := prometheus.NewRegistry()
	d := daemon.New(
		service_tester.NewOption(t),
		logger.New(t.Context()),
		memory.New(),
		y,
	)
	t.Cleanup(d.Close)
	assert.True(t, reaper.CountsOn(y))
}
