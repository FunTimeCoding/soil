package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"testing"
)

func NewService(
	t *testing.T,
	m soilFace.ReachableEmbedder,
	a face.Reranker,
) *service.Service {
	t.Helper()

	return New(t, NewOption(t, m, a), memory.New())
}
