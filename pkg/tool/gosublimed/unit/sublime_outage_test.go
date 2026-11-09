package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/sublime/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/unit/daemon_tester"
	"net/http"
	"testing"
)

func TestSublimeOutageIsAnsweredUnavailableWithoutCaptureOverRest(
	t *testing.T,
) {
	s := base.New(t)
	s.MockClient.Failure = refusedDial()
	c := restClientOf(t, s)

	for range 2 {
		answer, e := c.GetViewsWithResponse(t.Context())
		assert.FatalOnError(t, e)
		assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode())
		assert.FatalNotNil(t, answer.JSON503)
	}

	assert.Count(t, 0, s.Reporter.Events())
}

func TestSublimeOutageFailsToolWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.MockClient.Failure = refusedDial()
	c := s.Connect(t)
	defer c.Close()
	assert.StringContains(
		t,
		"connection refused",
		c.MustCallToolError(constant.ListViews, map[string]any{}),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestDaemonHandsOutSublimeGuard(t *testing.T) {
	_, okay := daemon_tester.New(t, mock_client.New()).Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestDaemonClosesClientItBuilt(t *testing.T) {
	c := mock_client.New()
	daemon_tester.New(t, c).Close()
	assert.Integer(t, 1, c.Closed())
}
