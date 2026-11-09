package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestClosingDaemonClosesEveryInstanceClient(t *testing.T) {
	d, clients := daemon_tester.New(
		t,
		memory.New(),
		constant.FixtureInstance,
		constant.FixtureOtherInstance,
	)
	d.Close()
	assert.Integer(t, 1, clients[constant.FixtureInstance].Closed)
	assert.Integer(t, 1, clients[constant.FixtureOtherInstance].Closed)
}

func TestDaemonHandsItsGuardsAndThenItsClockToLifecycle(t *testing.T) {
	d, _ := daemon_tester.New(
		t,
		memory.New(),
		constant.FixtureInstance,
		constant.FixtureOtherInstance,
	)
	assert.Count(t, 3, d.Workers())
}

func TestDaemonClosesItsClockAfterLifecycleDid(t *testing.T) {
	d, _ := daemon_tester.New(t, memory.New(), constant.FixtureInstance)
	w := d.Workers()
	w[len(w)-1].Close()
	d.Close()
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d, _ := daemon_tester.New(
		t,
		memory.New(),
		constant.FixtureInstance,
		constant.FixtureOtherInstance,
	)
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
