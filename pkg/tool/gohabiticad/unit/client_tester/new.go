package client_tester

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/client"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/base"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	b *base.Server,
	token string,
) *client.Client {
	t.Helper()
	m := http.NewServeMux()
	b.Daemon().Mount(
		b.Reporter,
		mock_recorder.New(),
		guard.New(m, []string{generativeConstant.ModelContextTestToken}),
	)
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)
	host, port := network.SplitHostPort(s.Listener.Addr().String())
	result := client.New(locator.New(host).Port(port).Insecure(), token)
	t.Cleanup(result.Close)

	return result
}
