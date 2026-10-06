package band

import (
	"github.com/funtimecoding/soil/pkg/band/response"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustPowerState() *response.PowerState {
	result, e := c.PowerState()
	errors.PanicOnError(e)

	return result
}
