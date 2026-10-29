package base

import (
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func serve(
	t *testing.T,
	configure func(*http.ServeMux),
	k *connector.Client,
) *Server {
	t.Helper()
	upstream := mattermost_client_tester.New(t, configure)
	r := memory.New()
	v := service_tester.Reported(t, upstream.Client, k, nil, r)

	return &Server{
		Upstream: upstream,
		Service:  v,
		Reporter: r,
		ContextServer: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				gomattermostd.Mount(v, r, mock_recorder.New(), g)
			},
		),
	}
}
