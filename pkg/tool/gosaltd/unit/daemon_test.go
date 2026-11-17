package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/base"
	"go.uber.org/goleak"
	"testing"
)

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
	d.Close()
	goleak.VerifyNone(t, before)
}

func TestDaemonCountsReapsOnItsRegistry(t *testing.T) {
	s := base.New(t)
	s.Daemon()
	assert.True(t, reaper.CountsOn(s.Registry))
}
