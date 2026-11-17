package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonHandsOneGuardPerHypervisorOutBeforePoller(t *testing.T) {
	s := base.New(t, service_tester.WithInstances("alfa", "bravo"))
	workers := s.Daemon().Workers()
	assert.Count(t, 3, workers)
	_, first := workers[0].(*peer.Guard)
	_, second := workers[1].(*peer.Guard)
	_, third := workers[2].(*peer.Guard)
	assert.True(t, first)
	assert.True(t, second)
	assert.False(t, third)
}

func TestDaemonHandsEveryGuardOutBeforeItsClocks(t *testing.T) {
	s := base.New(t, service_tester.WithNetbox())
	workers := s.Daemon().Workers()
	assert.Count(t, 4, workers)
	_, first := workers[0].(*peer.Guard)
	_, second := workers[1].(*peer.Guard)
	_, third := workers[2].(*peer.Guard)
	assert.True(t, first)
	assert.True(t, second)
	assert.False(t, third)
}

func TestClosingDaemonClosesEveryClientItBuilt(t *testing.T) {
	s := base.New(
		t,
		service_tester.WithInstances("alfa", "bravo"),
		service_tester.WithNetbox(),
	)
	s.Daemon().Close()
	assert.Integer(t, 1, s.Alfa().Hypervisor.Closed)
	assert.Integer(t, 1, s.Bravo().Hypervisor.Closed)
	assert.Integer(t, 1, s.Netbox.Closed)
}

func TestClosingDaemonClosesShellsItDialled(t *testing.T) {
	s := base.New(t)
	s.Tools().MustCallTool(constant.ListSnippets, nil)
	s.Daemon().Close()
	assert.Integer(t, 1, s.Alfa().Shell.Closed)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := base.New(t, service_tester.WithNetbox())
	before := goleak.IgnoreCurrent()
	d := s.Daemon()

	for _, w := range d.Workers() {
		w.Start()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}
