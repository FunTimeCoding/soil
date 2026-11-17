package client_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func NewReported(
	t *testing.T,
	j *mock_client.Client,
) (*client.Client, *memory.Memory) {
	t.Helper()
	o := service_tester.New(t, service_tester.WithJira(j))
	s := model_context_server.Mount(t, o.Daemon(), o.Reporter)
	result := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result, o.Reporter
}
