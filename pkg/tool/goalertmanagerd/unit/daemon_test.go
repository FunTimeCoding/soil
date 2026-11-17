package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonHandsTwoGuardsPerInstanceToLifecycle(t *testing.T) {
	s := base.New(t, service_tester.WithInstances("alfa", "bravo"))
	assert.Count(t, 4, s.Daemon().Workers())
}

func TestClosingDaemonClosesBothClientsOfEveryInstance(t *testing.T) {
	s := base.New(t)
	s.Daemon().Close()
	assert.Integer(t, 1, s.Alfa().Alertmanager.Closed)
	assert.Integer(t, 1, s.Alfa().Prometheus.Closed)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := base.New(t)
	before := goleak.IgnoreCurrent()
	d := s.Daemon()
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
