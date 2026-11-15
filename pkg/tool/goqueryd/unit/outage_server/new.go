package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_embedder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	m := mock_embedder.New()
	t.Cleanup(m.Close)
	m.Fail(outage_tester.RefusedDial())
	m.FailEmbedding(outage_tester.RefusedDial())
	k := mock_reranker.New()
	t.Cleanup(func() { errors.PanicClose(k) })
	d := daemon.New(
		service_tester.NewOption(t, m, k),
		logger.New(t.Context()),
		r,
		prometheus.NewRegistry(),
	)
	t.Cleanup(d.Close)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
	)
	t.Cleanup(result.Close)

	return result, outage_tester.Guard(
		t,
		d.Workers(),
		constant.EmbedderPeer,
	).Unavailable
}
