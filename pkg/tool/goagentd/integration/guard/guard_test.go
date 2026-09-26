package guard

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goagentd"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/runner"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			goagentd.Mount(runner.New(t.TempDir()), mock_recorder.New(), g)
		},
	)
	defer v.Stop()
	v.VerifyBase(t)
	v.VerifyInterface(t)
	v.VerifyGuarded(t, "/api/status")
}
