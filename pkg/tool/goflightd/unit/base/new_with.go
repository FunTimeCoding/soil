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
	"github.com/funtimecoding/soil/pkg/tool/goflightd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
	"time"
)

func NewWith(
	t *testing.T,
	clock func() time.Time,
) *Server {
	t.Helper()
	s := mock_system.New()
	t.Cleanup(s.Close)
	o := NewOption(t, s)
	o.Clock = clock
	l := logger.New(t.Context())
	r := memory.New()
	d := daemon.New(o, l, r)
	t.Cleanup(d.Close)
	m := lite.New(l, o.LitePath)
	t.Cleanup(func() { relational.CloseMapper(m) })
	h := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
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
		Mapper:   m,
		System:   s,
		Reporter: r,
		Client:   c,
		Server:   h,
	}
}
