package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/web"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func NewWithProject(
	t *testing.T,
	project []string,
) *Server {
	t.Helper()
	c := service_tester.NewConfluence(t)
	v := service_tester.NewService(t, service_tester.NewJira(t), c, project)
	k := worker.New(
		v,
		constant.PollInterval,
		logger.New(context.Background()),
		memory.New(),
	)
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goatlassiand.Mount(
				v,
				web.New(k),
				memory.New(),
				mock_recorder.New(),
				g,
			)
		},
	)
	result := &Server{MockClient: c, Server: s}
	t.Cleanup(result.Close)

	return result
}
