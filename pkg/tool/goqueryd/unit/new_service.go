package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func newService(
	t *testing.T,
	m face.ReachableEmbedder,
	a queryd.Reranker,
) *service.Service {
	t.Helper()

	return service_tester.New(
		t,
		service_tester.NewOption(t, m, a),
		memory.New(),
	)
}
