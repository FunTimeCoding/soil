package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestClosingDaemonClosesEveryInstanceClient(t *testing.T) {
	s := base.New(t, service_tester.WithInstances("alfa", "bravo"))
	s.Daemon().Close()
	assert.Integer(t, 1, s.Alfa().Closed)
	assert.Integer(t, 1, s.Bravo().Closed)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := base.New(t, service_tester.WithInstances("alfa", "bravo"))
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

func TestDaemonHandsOneGuardPerInstanceToLifecycle(t *testing.T) {
	s := base.New(t, service_tester.WithInstances("alfa", "bravo"))
	assert.Count(t, 2, s.Daemon().Workers())
}
