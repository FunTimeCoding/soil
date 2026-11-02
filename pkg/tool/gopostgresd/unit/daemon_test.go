package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestClosingTheDaemonClosesEveryInstanceClient(t *testing.T) {
	d, clients := daemon_tester.New(t, "alfa", "bravo")
	d.Close()
	assert.Integer(t, 1, clients["alfa"].Closed)
	assert.Integer(t, 1, clients["bravo"].Closed)
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d, _ := daemon_tester.New(t, "alfa", "bravo")
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

func TestTheDaemonHandsOneGuardPerInstanceToTheLifecycle(t *testing.T) {
	d, _ := daemon_tester.New(t, "alfa", "bravo")
	assert.Count(t, 2, d.Workers())
}
