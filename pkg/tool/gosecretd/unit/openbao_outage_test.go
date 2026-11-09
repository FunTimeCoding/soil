package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/unit/base"
	"net/http"
	"testing"
)

func TestOpenbaoOutageIsAnsweredUnavailableWithoutCaptureOverRest(
	t *testing.T,
) {
	s := base.New(t)
	defer s.Close()
	s.MockClient.Failure = refusedDial()
	c := interfaceClient(t, s)

	for range 2 {
		answer, e := c.ListSecretsWithResponse(
			t.Context(),
			&client.ListSecretsParams{},
		)
		assert.FatalOnError(t, e)
		assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode())
		assert.FatalNotNil(t, answer.JSON503)
	}

	assert.Count(t, 0, s.Reporter.Events())
}

func TestOpenbaoOutageFailsToolWithoutCapture(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	s.MockClient.Failure = refusedDial()
	c := s.ContextServer.Connect(t)
	defer c.Close()
	assert.StringContains(
		t,
		"connection refused",
		c.MustCallToolError(constant.ListSecrets, map[string]any{}),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestDaemonHandsOutOpenbaoGuardFirst(t *testing.T) {
	d := newOpenbaoDaemon(t)
	_, okay := d.Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestWaitAnswersOnceOpenbaoAnswers(t *testing.T) {
	assert.FatalOnError(t, newOpenbaoDaemon(t).Wait(t.Context()))
}
