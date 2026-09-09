package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/platform"

func (c *Client) CreatePlatform(_ string) (*platform.Platform, error) {
	return nil, nil
}
