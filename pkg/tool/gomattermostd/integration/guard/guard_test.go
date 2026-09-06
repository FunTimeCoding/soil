package guard

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	var c *mattermost.Client
	var m *monitor.Monitor
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			gomattermostd.Mount(
				c,
				m,
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
