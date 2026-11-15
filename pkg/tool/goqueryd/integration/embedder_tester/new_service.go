package embedder_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func NewService(t *testing.T) *service.Service {
	t.Helper()
	o := service_tester.NewOption(t, New(t), mock_reranker.New())
	c := NewCounter(t)
	o.EmbedCounter = func() face.EmbedCounter { return c }

	return service_tester.New(t, o, memory.New())
}
