package client_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	s face.OpnsenseSource,
) *client.Client {
	t.Helper()
	m := http.NewServeMux()
	gopnsensed.Mount(
		service_tester.New(t, s),
		memory.New(),
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
