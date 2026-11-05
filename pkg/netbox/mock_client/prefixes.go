package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/prefix"

func (c *Client) Prefixes() ([]*prefix.Prefix, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
