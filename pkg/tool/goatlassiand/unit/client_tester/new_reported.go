package client_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func NewReported(
	t *testing.T,
	j *mock_client.Client,
) (*client.Client, *memory.Memory) {
	t.Helper()
	o := service_tester.New(t, service_tester.WithJira(j))
	d := o.Daemon()
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(o.Reporter, mock_recorder.New(), g)
		},
	)
	result := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result, o.Reporter
}
