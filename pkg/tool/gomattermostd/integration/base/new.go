package base

import (
	"github.com/funtimecoding/soil/pkg/chat/integration/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/mock_indexer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	configure func(*http.ServeMux),
) *Server {
	t.Helper()
	upstream := mattermost_client_tester.New(t, configure)
	s := store.New(lite.NewMemory())
	t.Cleanup(s.Close)
	index := mock_indexer.New()

	return &Server{
		Upstream: upstream,
		Store:    s,
		Indexer:  index,
		ContextServer: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				gomattermostd.Mount(
					upstream.Client,
					nil,
					s,
					index,
					memory.New(),
					mock_recorder.New(),
					constant.DefaultVersion,
					g,
				)
			},
		),
	}
}
