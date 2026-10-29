package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/rejecting_host"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestARejectedHostIsAnsweredQuietlyOverModelContext(t *testing.T) {
	m := memory.New()
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			gopnsensed.Mount(
				service_tester.New(t, rejecting_host.New()),
				m,
				mock_recorder.New(),
				g,
			)
		},
	)
	defer s.Close()
	c := s.Connect(t)
	defer c.Close()
	assert.String(
		t,
		"host rejected: ip: invalid address",
		c.MustCallToolError(
			constant.AddHost,
			map[string]any{constant.ParameterHost: "alfa"},
		),
	)
	assert.Count(t, 0, m.Events())
}
