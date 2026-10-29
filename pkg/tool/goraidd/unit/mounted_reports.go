package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goraidd"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func mountedReports(
	t *testing.T,
	output string,
) *model_context_server.Server {
	t.Helper()
	v := service_tester.New(t, t.TempDir(), output)
	u := web.New(v, memory.New())
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goraidd.Mount(v, u, memory.New(), mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)

	return s
}
