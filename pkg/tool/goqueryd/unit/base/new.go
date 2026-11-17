package base

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/shared_reranker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	e := ollama.NewEnvironment()
	t.Cleanup(e.Close)
	k := shared_reranker.Lease()
	t.Cleanup(func() { errors.PanicClose(k) })
	s := service_tester.New(
		t,
		service_tester.WithEmbedder(e),
		service_tester.WithReranker(k),
	)
	d := s.Daemon()
	result := &Server{
		t:      t,
		daemon: d,
		Server: model_context_server.New(
			t,
			func(g *guard.Mux) {
				d.Mount(s.Reporter, mock_recorder.New(), g)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
