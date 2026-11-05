package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonClosesTheClientItBuilt(t *testing.T) {
	c := service_tester.NewSource(t)
	daemon.New(
		service_tester.NewOption(t, c),
		logger.New(t.Context()),
		memory.New(),
		metric.New(),
	).Close()
	assert.Integer(t, 1, c.Closed)
}

func TestTheIntentClockRunsOnlyWithPowerIntent(t *testing.T) {
	assert.Count(t, 1, newDaemon(t, service_tester.NewSource(t)).Workers())
	o := service_tester.NewOption(t, service_tester.NewSource(t))
	o.PowerIntent = true
	d := daemon.New(o, logger.New(t.Context()), memory.New(), metric.New())
	t.Cleanup(d.Close)
	assert.Count(t, 2, d.Workers())
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	o := service_tester.NewOption(t, service_tester.NewSource(t))
	o.PowerIntent = true
	before := goleak.IgnoreCurrent()
	d := daemon.New(o, logger.New(t.Context()), memory.New(), metric.New())
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
