package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/option"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestClosingDaemonClosesEveryEngine(t *testing.T) {
	d, engines := daemon_tester.New(t, "alfa", "bravo")
	d.Close()
	assert.Integer(t, 1, engines["alfa"].Closed)
	assert.Integer(t, 1, engines["bravo"].Closed)
}

func TestDaemonHandsOneGuardPerHostWithoutNetbox(t *testing.T) {
	d, _ := daemon_tester.New(t, "alfa", "bravo")
	assert.Count(t, 2, d.Workers())
}

func TestHostWithoutPasswordIsRefusedNotFatal(t *testing.T) {
	i := instance.Instance{Name: "alfa", Host: "alfa.example"}
	_, e := option.BuildEngine(&i)
	assert.True(t, validation.Is(e))
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
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
