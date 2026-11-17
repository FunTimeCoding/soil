package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestItermOutageFailsToolWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	c := s.Tools()

	for range 2 {
		assert.StringContains(
			t,
			"connection refused",
			c.MustCallToolError(constant.ListSessions, map[string]any{}),
		)
	}

	assert.Count(t, 0, s.Reporter.Events())
}

func TestDaemonHandsOutItermGuard(t *testing.T) {
	_, okay := base.New(t).Daemon().Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}
