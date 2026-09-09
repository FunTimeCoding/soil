package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"

func (c *Client) Lease(
	_ string,
	_ string,
) (*lease.Lease, error) {
	return nil, nil
}
