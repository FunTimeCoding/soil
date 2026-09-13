package guard

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/raid_parser"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goraidd"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"path/filepath"
	"testing"
)

func TestGuard(t *testing.T) {
	logCache := t.TempDir()
	system.WriteFile(
		filepath.Join(logCache, "LogDataCache.json"),
		[]byte("{}\n"),
		0o600,
	)
	s := store.New(
		lite.NewMemory(),
		logCache,
		filepath.Join(t.TempDir(), "elite"),
		logger.New(t.Context()),
		memory.New(),
	)
	p := raid_parser.New(locator.New("localhost"), "guard")
	authorization := client.New(
		"https://gate.example.org",
		"tester",
		"tester-secret",
		constant.SignInPath,
		"https://raid.example.org/callback",
		client.DeriveKey("tester-encryption-secret"),
	)
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goraidd.Mount(
				s,
				t.TempDir(),
				web.New(s, t.TempDir(), t.TempDir(), p, authorization),
				memory.New(),
				mock_recorder.New(),
				g,
			)
		},
	)
	defer v.Stop()
	v.VerifyBase(t)
	v.VerifyInterface(t)
	v.VerifyOpen(t, "/favicon.ico")
}
