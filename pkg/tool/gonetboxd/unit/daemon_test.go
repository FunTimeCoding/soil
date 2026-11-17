package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesClientItBuilt(t *testing.T) {
	o := service_tester.New(t)
	o.Daemon().Close()
	assert.Integer(t, 1, o.Source.Closed)
}

func TestIntentClockRunsOnlyWithPowerIntent(t *testing.T) {
	assert.Count(t, 1, service_tester.New(t).Daemon().Workers())
	d := service_tester.New(t, service_tester.WithPowerIntent()).Daemon()
	assert.Count(t, 2, d.Workers())
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	o := service_tester.New(t, service_tester.WithPowerIntent())
	before := goleak.IgnoreCurrent()
	d := o.Daemon()
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
