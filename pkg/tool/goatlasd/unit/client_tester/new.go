package client_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/client"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	v *service.Service,
	token string,
) *client.Client {
	t.Helper()
	m := http.NewServeMux()
	goatlasd.Mount(
		v,
		web.New(v),
		memory.New(),
		mock_recorder.New(),
		guard.New(m, []string{constant.ModelContextTestToken}),
	)
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)
	host, port := network.SplitHostPort(s.Listener.Addr().String())
	result := client.New(locator.New(host).Port(port).Insecure(), token)
	t.Cleanup(result.Close)

	return result
}
