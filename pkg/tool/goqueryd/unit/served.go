package unit

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"testing"
)

func served(
	t *testing.T,
	o *option.Query,
	r face.Reporter,
) *model_context_server.Server {
	t.Helper()
	d := daemon.New(o, logger.New(t.Context()), r, prometheus.NewRegistry())
	t.Cleanup(d.Close)
	result := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(result.Close)

	return result
}
