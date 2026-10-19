package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/rejecting_stat"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestAnInvalidStatIsAnsweredQuietlyOverModelContext(t *testing.T) {
	m := memory.New()
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			gohabiticad.Mount(rejecting_stat.New(), m, mock_recorder.New(), g)
		},
	)
	defer s.Close()
	c := s.Connect(t)
	defer c.Close()
	assert.String(
		t,
		`"bravo" is not a valid Stat.`,
		c.MustCallToolError(
			constant.AllocateStat,
			map[string]any{constant.Stat: "bravo"},
		),
	)
	assert.Count(t, 0, m.Events())
}
