package base

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	v := service.New(
		directory.New("directory.example.org", "dc=example,dc=org"),
	)
	authorization := client.New(
		"https://gate.example.org",
		"tester",
		"tester-secret",
		webConstant.SignInPath,
		"https://directory.example.org/callback",
		client.DeriveKey("tester-encryption-secret"),
	)

	return &Server{
		Service:       v,
		Authorization: authorization,
		Server: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				godirectoryd.Mount(
					v,
					web.New(v, authorization),
					memory.New(),
					mock_recorder.New(),
					library.DefaultVersion,
					g,
				)
			},
		),
	}
}
