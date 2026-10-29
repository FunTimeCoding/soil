package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gogated"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			gogated.Mount(
				s.Service,
				web.New(s.Service, s.Reporter),
				memory.New(),
				mock_recorder.New(),
				g,
			)
		},
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyOpen(t, "/jwks")
	v.VerifyOpenPost(t, "/token")
	v.VerifyOpen(t, "/favicon.ico")
	v.VerifyModelContext(t)
	v.VerifySignIn(t)
}
