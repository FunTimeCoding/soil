package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand"
	atlassianConstant "github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/mock_client/mock_jira"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/web"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	var j *jira.Client
	var f *confluence.Client
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goatlassiand.Mount(
				mock_jira.New(),
				c,
				web.New(
					worker.New(
						j,
						f,
						atlassianConstant.PollInterval,
						logger.New(context.Background()),
						memory.New(),
					),
				),
				memory.New(),
				mock_recorder.New(),
				constant.DefaultVersion,
				g,
			)
		},
	)

	return &Server{MockClient: c, Server: v}
}
