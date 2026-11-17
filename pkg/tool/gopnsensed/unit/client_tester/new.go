package client_tester

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/client"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/base"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	source face.OpnsenseClient,
	o ...base.Option,
) *client.Client {
	t.Helper()
	b := base.New(t, source, o...)
	m := http.NewServeMux()
	b.Daemon().Mount(
		b.Reporter,
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
