package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
	result := &Client{
		MachineSnapshots:   map[int64][]client.Snapshot{},
		ContainerSnapshots: map[int64][]client.Snapshot{},
		Snippets:           map[string]string{},
		Plan:               &client.ReconcilePlan{},
	}
	t.Cleanup(result.Close)

	return result
}
