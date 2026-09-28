package band

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustPowerState() *PowerStateResponse {
	result, e := c.PowerState()
	errors.PanicOnError(e)

	return result
}
