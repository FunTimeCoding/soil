package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/service_tester"
	"testing"
	"time"
)

func TestMasterOutageFailsToolWithoutCapture(t *testing.T) {
	s := base.New(t, service_tester.WithDownMaster())
	c := s.Served().Connect(t)
	defer c.Close()
	assert.StringContains(
		t,
		"connection refused",
		c.MustCallToolError(constant.Keys, map[string]any{}),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestKeySyncStaysQuietWhileMasterIsDown(t *testing.T) {
	s := base.New(
		t,
		service_tester.WithDownMaster(),
		service_tester.WithKeySyncInterval(20*time.Millisecond),
	)
	workers := s.Daemon().Workers()

	for _, w := range workers {
		w.Start()
	}

	time.Sleep(200 * time.Millisecond)

	for _, w := range workers {
		w.Close()
	}

	assert.Count(t, 0, s.Reporter.Events())
}

func TestDaemonHandsOutMasterGuardFirst(t *testing.T) {
	s := base.New(t, service_tester.WithDownMaster())
	_, okay := s.Daemon().Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}
