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
)

func (o *Tester) Reported(
	j *mock_client.Client,
) (*client.Client, *memory.Memory) {
	o.t.Helper()
	v := service_tester.New(o.t, service_tester.WithJira(j))
	s := model_context_server.Mount(o.t, v.Daemon(), v.Reporter)
	result := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	o.t.Cleanup(result.Close)

	return result, v.Reporter
}
