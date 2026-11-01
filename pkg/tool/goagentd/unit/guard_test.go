package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goagentd"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	v := service_tester.New(
		t,
		service_tester.Fixed(mock_stream.New(nil, false, nil)),
	)
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goagentd.Mount(v, memory.New(), mock_recorder.New(), g)
		},
	)
	defer s.Close()
	s.VerifyBase(t)
	s.VerifyGuarded(t, "/api/status")
	s.VerifyGuarded(t, "/api/run")
}
