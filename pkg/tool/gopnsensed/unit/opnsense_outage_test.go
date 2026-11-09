package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/opnsense/mock_client"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func TestOpnsenseOutageIsAnsweredUnavailableWithoutCaptureOverRest(
	t *testing.T,
) {
	c := mock_client.New()
	c.Failure = refusedDial()
	r := memory.New()
	s := client_tester.NewWith(t, c, r)

	for range 2 {
		_, e := s.ListLeases(t.Context(), "")
		assert.True(t, unavailable.Is(e))
	}

	assert.Count(t, 0, r.Events())
}

func TestOpnsenseOutageFailsToolWithoutCapture(t *testing.T) {
	c := mock_client.New()
	c.Failure = refusedDial()
	m := memory.New()
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			daemon_tester.New(t, c).Mount(m, mock_recorder.New(), g)
		},
	)
	defer s.Close()
	v := s.Connect(t)
	defer v.Close()
	assert.StringContains(
		t,
		"connection refused",
		v.MustCallToolError(constant.ListLeases, map[string]any{}),
	)
	assert.Count(t, 0, m.Events())
}

func TestDaemonHandsOutOpnsenseGuard(t *testing.T) {
	_, okay := daemon_tester.New(t, mock_client.New()).Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestWaitAnswersOnceOpnsenseAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		daemon_tester.New(t, mock_client.New()).Wait(t.Context()),
	)
}
