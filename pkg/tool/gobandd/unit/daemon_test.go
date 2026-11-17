package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/option"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestClosingDaemonClosesEveryEngine(t *testing.T) {
	s := base.New(t, service_tester.WithInstances("alfa", "bravo"))
	s.Daemon().Close()
	assert.Integer(t, 1, s.Alfa().Closed)
	assert.Integer(t, 1, s.Bravo().Closed)
}

func TestDaemonHandsOneGuardPerHostWithoutNetbox(t *testing.T) {
	s := base.New(t, service_tester.WithInstances("alfa", "bravo"))
	assert.Count(t, 2, s.Daemon().Workers())
}

func TestHostWithoutPasswordIsRefusedNotFatal(t *testing.T) {
	i := instance.Instance{Name: "alfa", Host: "alfa.example"}
	_, e := option.BuildEngine(&i)
	assert.True(t, validation.Is(e))
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
