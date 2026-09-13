package base

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_snippet"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/web"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"testing"
	"time"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	c.AddNode("test")
	s := mock_service.New("test", c, mock_snippet.New())
	k := worker.New(
		s,
		time.Hour,
		prometheus.NewRegistry(),
		logger.New(t.Context()),
		memory.New(),
	)
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goproxmoxd.Mount(
				s,
				web.New(s, k),
				memory.New(),
				mock_recorder.New(),
				constant.DefaultVersion,
				g,
			)
		},
	)

	return &Server{MockClient: c, Server: v}
}
