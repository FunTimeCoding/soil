package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	s := service_tester.New(t).Service
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			r := memory.New()
			gosproutd.Mount(s, web.New(s, r), r, mock_recorder.New(), g)
		},
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyOpen(t, constant.DashboardPath)
	v.VerifyOpen(t, constant.SessionsPath)
	v.VerifyOpen(t, constant.CountersPath)
	v.VerifyOpen(t, "/event")
	v.VerifyModelContext(t)
}
