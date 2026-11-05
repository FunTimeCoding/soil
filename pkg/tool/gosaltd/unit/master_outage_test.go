package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
	"time"
)

func TestAMasterOutageFailsTheToolWithoutACapture(t *testing.T) {
	r := memory.New()
	d := newDownDaemon(t, r)
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	defer s.Close()
	c := s.Connect(t)
	defer c.Close()
	assert.StringContains(
		t,
		"connection refused",
		c.MustCallToolError(constant.Keys, map[string]any{}),
	)
	assert.Count(t, 0, r.Events())
}

func TestTheKeySyncStaysQuietWhileTheMasterIsDown(t *testing.T) {
	r := memory.New()
	d := newDownDaemon(t, r)
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	time.Sleep(200 * time.Millisecond)

	for _, w := range workers {
		w.Close()
	}

	assert.Count(t, 0, r.Events())
}

func TestTheDaemonHandsOutTheMasterGuardFirst(t *testing.T) {
	_, okay := newDownDaemon(t, memory.New()).Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}
