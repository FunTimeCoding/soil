package base

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.New(lite.NewMemory())
	f := mock_client.New()
	v := service.New(
		s,
		publish.New(
			f,
			constant.FixtureProject,
			constant.PublishBranch,
			constant.FixtureAuthorityDirectory,
			constant.FixtureClusterAuthority,
			constant.FixtureSecretPath,
		),
	)
	r := memory.New()
	authorization := client.New(
		"https://gate.example.org",
		"tester",
		"tester-secret",
		webConstant.SignInPath,
		"https://certificate.example.org/callback",
		client.DeriveKey("tester-encryption-secret"),
	)

	return &Server{
		Store:         s,
		Service:       v,
		Forge:         f,
		Authorization: authorization,
		Server: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				gocertificated.Mount(
					s,
					v,
					web.New(s, v, authorization),
					r,
					mock_recorder.New(),
					library.DefaultVersion,
					g,
				)
			},
		),
	}
}
