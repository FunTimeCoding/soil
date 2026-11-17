package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/opnsense/mock_client"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOpnsenseOutageIsAnsweredUnavailableWithoutCaptureOverRest(
	t *testing.T,
) {
	c := mock_client.New(t)
	c.Failure = outage_tester.RefusedDial()
	r := memory.New()
	s := client_tester.New(t, c, base.WithReporter(r))

	for range 2 {
		_, e := s.ListLeases(t.Context(), "")
		assert.True(t, unavailable.Is(e))
	}

	assert.Count(t, 0, r.Events())
}

func TestOpnsenseOutageFailsToolWithoutCapture(t *testing.T) {
	c := mock_client.New(t)
	c.Failure = outage_tester.RefusedDial()
	b := base.New(t, c)
	v := b.Tools()
	assert.StringContains(
		t,
		"connection refused",
		v.MustCallToolError(constant.ListLeases, map[string]any{}),
	)
	assert.Count(t, 0, b.Reporter.Events())
}

func TestDaemonHandsOutOpnsenseGuard(t *testing.T) {
	_, okay := base.New(t, mock_client.New(t)).Daemon().Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestWaitAnswersOnceOpnsenseAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		base.New(t, mock_client.New(t)).Daemon().Wait(t.Context()),
	)
}
