package client_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client"
	"github.com/funtimecoding/soil/pkg/atlassian/face"
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
	j face.JiraClient,
) (*client.Client, *memory.Memory) {
	t.Helper()
	r := memory.New()
	d := service_tester.NewDaemon(
		t,
		service_tester.NewOption(j, mock_client.New(), nil),
		r,
	)
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	result := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result, r
}
