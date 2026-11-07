package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/openbao/face"
	"github.com/funtimecoding/soil/pkg/openbao/mock_client"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	o := option.New()
	o.Source = func() face.OpenbaoClient { return c }
	r := memory.New()
	d := daemon.New(
		o,
		logger.New(context.Background()),
		r,
		prometheus.NewRegistry(),
	)
	t.Cleanup(d.Close)

	return &Server{
		MockClient: c,
		Reporter:   r,
		ContextServer: model_context_server.New(
			t,
			func(g *guard.Mux) {
				d.Mount(r, mock_recorder.New(), g)
			},
		),
	}
}
