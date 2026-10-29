package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gochromed"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			gochromed.Mount(
				service_tester.New(t, mock_client.New()),
				memory.New(),
				mock_recorder.New(),
				g,
			)
		},
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
