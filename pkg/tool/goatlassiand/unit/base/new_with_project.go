package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func NewWithProject(
	t *testing.T,
	project []string,
) *Server {
	t.Helper()
	j := service_tester.NewJira(t)
	c := service_tester.NewConfluence(t)
	r := memory.New()
	d := service_tester.NewDaemon(t, service_tester.NewOption(j, c, project), r)
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	result := &Server{MockClient: c, Jira: j, Daemon: d, Server: s}
	t.Cleanup(result.Close)

	return result
}
