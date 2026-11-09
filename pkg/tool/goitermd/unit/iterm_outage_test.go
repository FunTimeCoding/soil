package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/iterm/mock_client"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/daemon_tester"
	"testing"
)

func TestItermOutageFailsToolWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.MockClient.Failure = refusedDial()
	c := s.Connect(t)
	defer c.Close()

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
	_, okay := daemon_tester.New(t, mock_client.New()).Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}
