package technitium

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustCreateZone(
	name string,
	zoneType string,
) string {
	result, e := c.CreateZone(name, zoneType)
	errors.PanicOnError(e)

	return result
}
