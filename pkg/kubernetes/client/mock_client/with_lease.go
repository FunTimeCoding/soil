package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"

func (c *Client) WithLease(l *lease.Lease) *Client {
	c.lease = l

	return c
}
