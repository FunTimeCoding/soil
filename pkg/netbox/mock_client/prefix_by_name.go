package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/prefix"

func (c *Client) PrefixByName(_ string) (*prefix.Prefix, error) {
	return nil, c.Failure
}
