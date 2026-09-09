package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_disk"

func (c *Client) CreateVirtualDisk(
	_ string,
	_ string,
	_ int32,
) (*virtual_disk.Disk, error) {
	return nil, nil
}
