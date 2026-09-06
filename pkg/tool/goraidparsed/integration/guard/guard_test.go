package guard

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/server"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			goraidparsed.Mount(
				server.New(
					t.TempDir(),
					t.TempDir(),
					t.TempDir(),
					logger.New(t.Context()),
					memory.New(),
				),
				mock_recorder.New(),
				g,
			)
		},
	)
	defer v.Stop()
	v.VerifyBase(t)
	v.VerifyInterface(t)
}
