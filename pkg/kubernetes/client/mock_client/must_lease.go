package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"

func (c *Client) MustLease(
	_ string,
	_ string,
) *lease.Lease {
	return c.lease
}
