package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	r := memory.New()
	v := service_tester.New(
		t,
		service_tester.NewOption(
			t,
			service_tester.FixtureBoard(),
			newPrometheus(t),
			nil,
			nil,
		),
		r,
	)
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			godashboardd.Mount(web.New(v, r), g)
		},
	)
	defer s.Close()
	s.VerifyBase(t)
	s.VerifyOpen(t, "/favicon.ico")
	s.VerifySignIn(t)
}
