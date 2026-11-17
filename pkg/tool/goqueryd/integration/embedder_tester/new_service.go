package embedder_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func NewService(t *testing.T) *service.Service {
	t.Helper()

	return service_tester.New(
		t,
		service_tester.WithEmbedder(New(t)),
		service_tester.WithEmbedCounter(NewCounter(t)),
	).Service()
}
