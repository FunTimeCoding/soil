package base

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/shared_reranker"
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
		Server: model_context_server.Mount(t, d, s.Reporter),
	}
	t.Cleanup(result.Close)

	return result
}
