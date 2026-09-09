package mock_client

import "github.com/funtimecoding/soil/pkg/opnsense/pool"

func (c *Client) Pools(_ string) ([]*pool.Pool, error) {
	return nil, nil
}
