package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"path/filepath"
	"testing"
	"time"
)

func serve(
	t *testing.T,
	hiddenTag string,
	root string,
	clock func() time.Time,
) *Server {
	t.Helper()
	l := logger.New(t.Context())
	i := mock_client.New()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "memory.db")
	o.Query = func() face.Client { return i }
	o.HiddenTag = hiddenTag
	o.ReferenceRoot = root
	o.Clock = clock
	r := memory.New()
	v := daemon.New(o, l, r)
	d := connection.New(l, o.LitePath)
	result := &Server{
		t:        t,
		database: d,
		store:    store.New(d),
		daemon:   v,
		indexer:  i,
		Reporter: r,
		Server: model_context_server.New(
			t,
			func(g *guard.Mux) {
				v.Mount(r, mock_recorder.New(), g)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
