package mock_source

import "github.com/funtimecoding/soil/pkg/opnsense/lease"

func (c *Client) Leases(_ string) ([]*lease.Lease, error) {
	return nil, nil
}
