package client_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	s face.NetBoxClient,
) *client.Client {
	t.Helper()
	d := daemon.New(
		service_tester.NewOption(t, s),
		logger.New(t.Context()),
		memory.New(),
		metric.New(),
	)
	t.Cleanup(d.Close)
	m := http.NewServeMux()
	d.Mount(
		memory.New(),
		mock_recorder.New(),
		guard.New(m, []string{constant.ModelContextTestToken}),
	)
	h := httptest.NewServer(m)
	t.Cleanup(h.Close)
	host, port := network.SplitHostPort(h.Listener.Addr().String())
	result := client.New(
		locator.New(host).Port(port).Insecure(),
		constant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
