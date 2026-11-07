package base

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"path/filepath"
	"testing"
)

func serve(
	t *testing.T,
	configure func(*http.ServeMux),
	k face.Client,
) *Server {
	t.Helper()
	upstream := mattermost_client_tester.New(t, configure)
	r := memory.New()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "mattermost.db")
	o.Mattermost = func() chat.MattermostClient { return upstream.Client }
	o.Connector = func() face.Client { return k }
	d := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(d.Close)

	return &Server{
		Upstream: upstream,
		Store:    worker_tester.NewStoreOnFile(t, o.LitePath),
		Reporter: r,
		ContextServer: model_context_server.New(
			t,
			func(g *guard.Mux) {
				d.Mount(r, mock_recorder.New(), g)
			},
		),
	}
}
