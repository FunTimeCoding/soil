package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	q := service_tester.New(t, service_tester.WithReporter(r))
	q.Embedder.Fail(outage_tester.RefusedDial())
	q.Embedder.FailEmbedding(outage_tester.RefusedDial())
	d := q.Daemon()
	result := model_context_server.Mount(t, d, r)

	return result, outage_tester.New(t).Guard(
		d.Workers(),
		constant.EmbedderPeer,
	).Unavailable
}
