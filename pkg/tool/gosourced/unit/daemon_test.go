package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonHandsItsMemoryGuardToLifecycle(t *testing.T) {
	workers := base.New(t).Daemon().Workers()
	assert.Count(t, 1, workers)
	_, okay := workers[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestClosingDaemonClosesMemoryClientOnce(t *testing.T) {
	s := base.New(t)
	d := s.Daemon()
	d.Close()
	d.Close()
	assert.Integer(t, 1, s.Memory.Closed)
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
