package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/reachable_skip"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/shared_reranker"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	reachable_skip.Skip(t)
	e := ollama.NewEnvironment()
	t.Cleanup(e.Close)
	k := shared_reranker.Lease()
	t.Cleanup(func() { errors.PanicClose(k) })

	return &Tester{
		t: t,
		Service: service_tester.New(
			t,
			service_tester.WithEmbedder(e),
			service_tester.WithReranker(k),
		).Service(),
	}
}
