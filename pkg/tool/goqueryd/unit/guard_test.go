package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	v := newService(t, newEmbedder(t), newReranker(t))
	r := memory.New()
	w := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goqueryd.Mount(v, web.New(v, r), r, mock_recorder.New(), g)
		},
	)
	defer w.Close()
	w.VerifyBase(t)
	w.VerifyGuarded(t, "/api/status")
	w.VerifyOpen(t, constant.DashboardPath)
	w.VerifyModelContext(t)
}
