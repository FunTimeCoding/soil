package client_tester

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/client"
	generated "github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	j face.Client,
) *client.Client {
	t.Helper()
	m := http.NewServeMux()
	guard.New(m, []string{constant.ModelContextTestToken}).TokenMount(
		webConstant.InterfacePath,
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(
					service_tester.NewService(t, j, mock_client.New(), nil),
					memory.New(),
				),
				nil,
			),
			http.NewServeMux(),
		),
	)
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)
	host, port := network.SplitHostPort(s.Listener.Addr().String())
	result := client.New(
		locator.New(host).Port(port).Insecure(),
		constant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
