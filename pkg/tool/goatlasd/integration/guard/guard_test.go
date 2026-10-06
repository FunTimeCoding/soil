package guard

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	m := lite.NewMemory()
	migrate.AutoMigrate(m)
	s := store.New(m)
	authorization := client.New(
		"https://gate.example.org",
		"tester",
		"tester-secret",
		webConstant.SignInPath,
		"https://atlas.example.org/callback",
		client.DeriveKey("tester-encryption-secret"),
	)
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			goatlasd.Mount(
				s,
				web.New(s, authorization),
				reporter.New(constant.Identity.Name()),
				mock_recorder.New(),
				g,
			)
		},
	)
	defer v.Stop()
	v.VerifyBase(t)
	v.VerifyInterface(t)
	v.VerifyGuarded(t, "/api/placements")
	v.VerifyGuarded(t, "/api/places")
	v.VerifyGuarded(t, "/api/sightings")
	v.VerifyOpen(t, "/favicon.ico")
}
