package client_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/client"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	v *service_tester.Tester,
	token string,
) *client.Client {
	t.Helper()
	r := memory.New()
	d := daemon.New(
		v.Option,
		logger.New(t.Context()),
		r,
		prometheus.NewRegistry(),
	)
	t.Cleanup(d.Close)
	s := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	result := client.New(
		locator.New(webConstant.Localhost).Port(s.Port).Insecure(),
		token,
	)
	t.Cleanup(result.Close)

	return result
}
