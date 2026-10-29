package client_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/client"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/service"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(t *testing.T) *client.Client {
	t.Helper()
	m := http.NewServeMux()
	gooutpostd.Mount(
		service.New(),
		mock_recorder.New(),
		guard.New(m, []string{constant.ModelContextTestToken}),
	)
	v := httptest.NewServer(m)
	t.Cleanup(v.Close)
	host, port := network.SplitHostPort(v.Listener.Addr().String())
	result := client.New(
		locator.New(host).Port(port).Insecure(),
		constant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
