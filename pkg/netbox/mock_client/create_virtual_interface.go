package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateVirtualInterface(
	_ *virtual_machine.Machine,
	_ string,
) (*netbox.VMInterface, error) {
	return nil, nil
}
