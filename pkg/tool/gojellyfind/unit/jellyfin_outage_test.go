package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/base"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"net/http"
	"testing"
)

func TestJellyfinOutageIsAnsweredUnavailableWithoutCaptureOverRest(
	t *testing.T,
) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	c := s.GeneratedClient()

	for range 2 {
		answer, e := c.ListLibrariesWithResponse(t.Context())
		assert.FatalOnError(t, e)
		assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode())
		assert.FatalNotNil(t, answer.JSON503)
	}

	assert.Count(t, 0, s.Reporter.Events())
}

func TestJellyfinOutageFailsToolWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	c := s.Tools()
	assert.StringContains(
		t,
		"connection refused",
		c.MustCallToolError(constant.ListLibraries, map[string]any{}),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestDaemonHandsOutJellyfinGuard(t *testing.T) {
	_, okay := base.New(t).Daemon().Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestWaitAnswersOnceJellyfinAnswers(t *testing.T) {
	assert.FatalOnError(t, base.New(t).Daemon().Wait(t.Context()))
}
