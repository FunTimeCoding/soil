package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

func (c *Client) CreateVirtualMachine(
	_ string,
	_ *cluster.Cluster,
) (*virtual_machine.Machine, error) {
	return nil, nil
}
