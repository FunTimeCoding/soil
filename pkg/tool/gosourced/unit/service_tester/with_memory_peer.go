package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"testing"
)

func WithMemoryPeer(
	t *testing.T,
	c face.Client,
) (*Services, *memory.Memory) {
	t.Helper()
	r := memory.New()
	result := fromParts(t, inventory.New(), c, r)
	result.Workspace.UseIndex(IndexDirectory())

	return result, r
}
