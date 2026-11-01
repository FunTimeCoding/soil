package embedder

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func newService(t *testing.T) *service.Service {
	t.Helper()
	o := service_tester.NewOption(t, newEmbedder(t), mock_reranker.New())
	c := newCounter(t)
	o.EmbedCounter = func() face.EmbedCounter { return c }

	return service_tester.New(t, o, memory.New())
}
