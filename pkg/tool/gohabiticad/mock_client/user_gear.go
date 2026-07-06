package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/gear"

func (c *Client) UserGear() (*gear.Gear, error) {
	return c.gear, nil
}
