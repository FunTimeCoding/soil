package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/platform"

func (c *Client) PlatformByName(_ string) (*platform.Platform, error) {
	return nil, c.Failure
}
