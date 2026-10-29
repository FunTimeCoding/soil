package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	s := service_tester.New(t)
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			goatlasd.Mount(
				s.Service,
				web.New(s.Service),
				memory.New(),
				mock_recorder.New(),
				g,
			)
		},
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyGuarded(t, "/api/placements")
	v.VerifyGuarded(t, "/api/places")
	v.VerifyGuarded(t, "/api/sightings")
	v.VerifyOpen(t, "/favicon.ico")
	v.VerifySignIn(t)
}
