package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goflightd"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/service"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := mock_system.New()
	t.Cleanup(s.Close)
	o := NewOption(t, s)
	l := logger.New(t.Context())
	r := memory.New()
	v := service.New(o, l)
	t.Cleanup(v.Close)
	m := lite.New(l, o.LitePath)
	t.Cleanup(func() { relational.CloseMapper(m) })
	h := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goflightd.Mount(v, r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(h.Close)
	c, e := client.NewClientWithResponses(
		h.Address(webConstant.RootPath),
		client.WithHTTPClient(h.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)

	return &Server{
		Store:    store.New(m),
		Service:  v,
		System:   s,
		Reporter: r,
		Client:   c,
		Server:   h,
	}
}
