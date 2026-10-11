package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gogated"
	"github.com/funtimecoding/soil/pkg/tool/gogated/server"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	s := base.New(t)
	authorization := client.New(
		"https://gate.example.org",
		"tester",
		"tester-secret",
		constant.SignInPath,
		"https://gate.example.org/callback",
		client.DeriveKey("tester-encryption-secret"),
	)
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			gogated.Mount(
				server.New(s.Service),
				web.New(s.Service, authorization, "admin@example.org"),
				s.Service,
				memory.New(),
				mock_recorder.New(),
				g,
			)
		},
	)
	defer v.Stop()
	v.VerifyBase(t)
	v.VerifyOpen(t, "/jwks")
	v.VerifyOpenPost(t, "/token")
	v.VerifyOpen(t, "/favicon.ico")
	v.VerifyModelContext(t)
}
