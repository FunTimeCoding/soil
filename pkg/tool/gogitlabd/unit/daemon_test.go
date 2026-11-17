package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestClosingDaemonClosesEveryInstanceClient(t *testing.T) {
	s := base.New(
		t,
		service_tester.WithInstances(
			constant.FixtureInstance,
			constant.FixtureOtherInstance,
		),
	)
	s.Daemon().Close()
	assert.Integer(t, 1, s.Alfa().Closed)
	assert.Integer(t, 1, s.Bravo().Closed)
}

func TestDaemonHandsItsGuardsAndThenItsClockToLifecycle(t *testing.T) {
	s := base.New(
		t,
		service_tester.WithInstances(
			constant.FixtureInstance,
			constant.FixtureOtherInstance,
		),
	)
	assert.Count(t, 3, s.Daemon().Workers())
}

func TestDaemonClosesItsClockAfterLifecycleDid(t *testing.T) {
	d := base.New(t).Daemon()
	w := d.Workers()
	w[len(w)-1].Close()
	d.Close()
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := base.New(
		t,
		service_tester.WithInstances(
			constant.FixtureInstance,
			constant.FixtureOtherInstance,
		),
	)
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
