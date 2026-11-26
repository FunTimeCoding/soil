package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := base.New(t)
	before := goleak.IgnoreCurrent()
	s.Daemon().Close()
	goleak.VerifyNone(t, before)
}

func TestDaemonClosesClient(t *testing.T) {
	s := base.New(t)
	s.Daemon().Close()
	assert.True(t, s.Client.Closed())
}
