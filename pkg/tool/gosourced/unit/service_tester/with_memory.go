package service_tester

import (
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"testing"
)

func WithMemory(t *testing.T) (*Services, *mock_client.Client) {
	t.Helper()
	m := mock_client.New()
	t.Cleanup(m.Close)
	result := fromParts(t, inventory.New(), m)
	result.Workspace.UseIndex(IndexDirectory())

	return result, m
}
