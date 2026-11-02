package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func mountedReports(
	t *testing.T,
	output string,
) *model_context_server.Server {
	t.Helper()
	r := memory.New()
	d := service_tester.NewDaemon(
		t,
		service_tester.NewOption(
			t,
			t.TempDir(),
			output,
			service_tester.NewParser(t),
		),
		r,
	)
	s := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)

	return s
}
