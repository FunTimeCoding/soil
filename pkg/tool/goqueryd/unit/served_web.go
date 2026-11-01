package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func servedWeb(t *testing.T) (*model_context_server.Server, *service.Service) {
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

	return s, v
}
