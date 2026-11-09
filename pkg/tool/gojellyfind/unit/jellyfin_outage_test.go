package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/jellyfin/mock_client"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/daemon_tester"
	"net/http"
	"testing"
)

func TestJellyfinOutageIsAnsweredUnavailableWithoutCaptureOverRest(
	t *testing.T,
) {
	c := mock_client.New()
	c.Failure = refusedDial()
	r := memory.New()
	s := restClient(t, c, r)

	for range 2 {
		answer, e := s.ListLibrariesWithResponse(t.Context())
		assert.FatalOnError(t, e)
		assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode())
		assert.FatalNotNil(t, answer.JSON503)
	}

	assert.Count(t, 0, r.Events())
}

func TestJellyfinOutageFailsToolWithoutCapture(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	s.MockClient.Failure = refusedDial()
	c := s.ContextServer.Connect(t)
	defer c.Close()
	assert.StringContains(
		t,
		"connection refused",
		c.MustCallToolError(constant.ListLibraries, map[string]any{}),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestDaemonHandsOutJellyfinGuard(t *testing.T) {
	_, okay := daemon_tester.New(t, mock_client.New()).Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestWaitAnswersOnceJellyfinAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		daemon_tester.New(t, mock_client.New()).Wait(t.Context()),
	)
}
