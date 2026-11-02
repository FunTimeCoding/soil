package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"

func New() *Client {
	return &Client{
		MachineSnapshots:   map[int64][]client.Snapshot{},
		ContainerSnapshots: map[int64][]client.Snapshot{},
		Snippets:           map[string]string{},
		Plan:               &client.ReconcilePlan{},
	}
}
