package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/reachable_skip"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/shared_reranker"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	reachable_skip.Skip(t)

	return &Tester{
		t: t,
		Service: service_tester.New(
			t,
			service_tester.NewOption(
				t,
				ollama.NewEnvironment(),
				shared_reranker.Lease(),
			),
			memory.New(),
		),
	}
}
