package band

import (
	"github.com/funtimecoding/soil/pkg/band/response"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustScreenSettings() *response.ScreenSettings {
	result, e := c.ScreenSettings()
	errors.PanicOnError(e)

	return result
}
