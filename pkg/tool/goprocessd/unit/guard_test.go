package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/service"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"path/filepath"
	"testing"
)

func TestGuard(t *testing.T) {
	o := option.New()
	o.ProcfilePath = writeProcfile(t, "alfa: sleep 60\n")
	o.EnvrcPath = filepath.Join(filepath.Dir(o.ProcfilePath), ".envrc")
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goprocessd.Mount(
				service.New(o),
				memory.New(),
				mock_recorder.New(),
				g,
			)
		},
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
	v.VerifyGuarded(t, "/api/processes")
	v.VerifyGuarded(t, "/api/processes/alfa/log")
}
