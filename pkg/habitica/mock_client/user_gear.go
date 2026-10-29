package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/gear"

func (c *Client) UserGear() (*gear.Gear, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.gear, nil
}
