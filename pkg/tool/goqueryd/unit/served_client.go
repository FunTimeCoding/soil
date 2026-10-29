package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"testing"
)

func servedClient(t *testing.T) *client.Client {
	t.Helper()
	r := memory.New()
	v := service_tester.New(
		t,
		service_tester.NewOption(t, newEmbedder(t), newReranker(t)),
		r,
	)
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goqueryd.Mount(v, web.New(v, r), r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	result := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port),
		constant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
