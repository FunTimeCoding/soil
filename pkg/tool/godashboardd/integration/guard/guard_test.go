package guard

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"path/filepath"
	"testing"
)

func TestGuard(t *testing.T) {
	path := filepath.Join(t.TempDir(), "board.yaml")
	system.WriteFile(path, []byte("{}\n"), 0o600)
	c := store.New(lite.NewMemory())
	b := board.Load(path)
	m := mock_client.New()
	v := service.New(b, m, m, m, notifier.New(), logger.New(t.Context()))
	authorization := client.New(
		"https://gate.example.org",
		"tester",
		"tester-secret",
		constant.SignInPath,
		"https://dashboard.example.org/callback",
		client.DeriveKey("tester-encryption-secret"),
	)
	s := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			godashboardd.Mount(web.New(b, v, c, authorization), g)
		},
	)
	defer s.Stop()
	s.VerifyBase(t)
	s.VerifyOpen(t, "/favicon.ico")
}
