package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesWhatItBuiltOnce(t *testing.T) {
	s := service_tester.New(t, service_tester.WithUsage())
	d := s.Daemon()
	d.Close()
	d.Close()
	assert.Integer(t, 1, s.Prometheus.Closed)
	assert.Integer(t, 1, s.Usage.Closed)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d := service_tester.New(t, service_tester.WithUsage()).Daemon()
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

func TestDaemonHandsOutItsGuardsThenItsClock(t *testing.T) {
	d := service_tester.New(t, service_tester.WithUsage()).Daemon()
	assert.Count(t, 3, d.Workers())
}
