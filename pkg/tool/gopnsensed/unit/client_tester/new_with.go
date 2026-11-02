package client_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewWith(
	t *testing.T,
	s face.OpnsenseSource,
	r *memory.Memory,
) *client.Client {
	t.Helper()
	m := http.NewServeMux()
	daemon_tester.New(t, s).Mount(
		r,
		mock_recorder.New(),
		guard.New(m, []string{generativeConstant.ModelContextTestToken}),
	)
	v := httptest.NewServer(m)
	t.Cleanup(v.Close)
	host, port := network.SplitHostPort(v.Listener.Addr().String())
	result := client.New(
		locator.New(host).Port(port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
