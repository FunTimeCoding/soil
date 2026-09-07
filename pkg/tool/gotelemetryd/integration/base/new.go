package base

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.New(lite.NewMemory())
	r := memory.New()
	recorder := mock_recorder.New()
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			gotelemetryd.Mount(
				s,
				web.New(s),
				r,
				recorder,
				constant.DefaultVersion,
				g,
			)
		},
	)

	return &Server{Store: s, Recorder: recorder, Server: v}
}
