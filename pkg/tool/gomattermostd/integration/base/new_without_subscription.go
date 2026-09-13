package base

import (
	"github.com/funtimecoding/soil/pkg/chat/integration/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func NewWithoutSubscription(
	t *testing.T,
	configure func(*http.ServeMux),
) *Server {
	t.Helper()
	upstream := mattermost_client_tester.New(t, configure)

	return &Server{
		Upstream: upstream,
		ContextServer: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				gomattermostd.Mount(
					upstream.Client,
					nil,
					nil,
					nil,
					memory.New(),
					mock_recorder.New(),
					constant.DefaultVersion,
					g,
				)
			},
		),
	}
}
