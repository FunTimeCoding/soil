package guard

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			gopostgresd.Mount(
				service.New(inventory.New()),
				memory.New(),
				mock_recorder.New(),
				constant.DefaultVersion,
				g,
			)
		},
	)
	defer v.Stop()
	v.VerifyBase(t)
	v.VerifyInterface(t)
	v.VerifyModelContext(t)
}
