package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestAnInvalidStatIsAnsweredQuietlyOverModelContext(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = validation.New("%q is not a valid Stat.", "bravo")
	v := service_tester.Service(t, c)
	m := memory.New()
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			gohabiticad.Mount(v, m, mock_recorder.New(), g)
		},
	)
	defer s.Close()
	x := s.Connect(t)
	defer x.Close()
	assert.String(
		t,
		`"bravo" is not a valid Stat.`,
		x.MustCallToolError(
			constant.AllocateStat,
			map[string]any{constant.Stat: "bravo"},
		),
	)
	assert.Count(t, 0, m.Events())
}
