package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func TestAnUnknownLinkIsAnsweredQuietlyOverModelContext(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = not_found.Format("resource not found: links/%d", 12)
	m := memory.New()
	d := daemon_tester.New(t, c, m)
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(m, mock_recorder.New(), g)
		},
	)
	defer s.Close()
	x := s.Connect(t)
	defer x.Close()
	assert.String(
		t,
		"resource not found: links/12",
		x.MustCallToolError(
			constant.Delete,
			map[string]any{"type": "link", "identifier": 12},
		),
	)
	assert.Count(t, 0, m.Events())
}
