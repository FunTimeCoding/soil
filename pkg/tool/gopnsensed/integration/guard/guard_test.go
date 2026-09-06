package guard

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/opnsense"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	var c *opnsense.Client
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			gopnsensed.Mount(
				c,
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
