package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	l := logger.New(t.Context())
	i := mock_client.New()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "memory.db")
	o.Query = func() face.Client { return i }
	r := memory.New()
	v := service.New(o, l, r)
	d := connection.New(l, o.LitePath)
	result := &Server{
		t:        t,
		database: d,
		store:    store.New(d),
		service:  v,
		indexer:  i,
		Server: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				gomemoryd.Mount(v, web.New(v, r), r, mock_recorder.New(), g)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
