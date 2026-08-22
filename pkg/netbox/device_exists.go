package netbox

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) DeviceExists(name string) (bool, error) {
	_, e := c.DeviceByName(name)

	if e != nil {
		if not_found.Is(e) {
			return false, nil
		}

		return false, e
	}

	return true, nil
}
