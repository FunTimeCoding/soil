package habitica

import "github.com/funtimecoding/soil/pkg/habitica/gear"

func (c *Client) UserGear() (*gear.Gear, error) {
	result, e := c.user()

	if e != nil {
		return nil, e
	}

	return result.Items.Gear, nil
}
