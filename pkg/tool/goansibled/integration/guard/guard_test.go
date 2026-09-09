package guard

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goansibled"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/option"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/runner"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	s := store.New(lite.NewMemory(), "playbook_runs")
	n := runner.New(
		option.New(),
		s,
		logger.New(t.Context()),
		memory.New(),
		reaper.New(memory.New()),
	)
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			goansibled.Mount(
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
