package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gogated/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func TestGuard(t *testing.T) {
	d := daemon.New(base.NewOption(t), logger.New(t.Context()), memory.New())
	t.Cleanup(d.Close)
	v := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyOpen(t, "/jwks")
	v.VerifyOpenPost(t, "/token")
	v.VerifyFavicon(t)
	v.VerifyModelContext(t)
	v.VerifySignIn(t)
	v.VerifySession(t, webConstant.PalettePath)
}
