package base

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/chat/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
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
	result := &Server{
		Upstream:      upstream,
		Store:         worker_tester.New(t).StoreOnFile(o.LitePath),
		Reporter:      r,
		ContextServer: model_context_server.Mount(t, d, r),
		t:             t,
	}
	t.Cleanup(result.Close)

	return result
}
