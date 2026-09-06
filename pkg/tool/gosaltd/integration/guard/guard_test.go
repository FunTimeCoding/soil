package guard

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	var n *runner.Runner
	var s *store.Store
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			gosaltd.Mount(
				n,
				s,
				memory.New(),
				mock_recorder.New(),
				constant.DefaultVersion,
				g,
			)
		},
	)
	defer v.Stop()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
