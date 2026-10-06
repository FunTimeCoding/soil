package band

import (
	"github.com/funtimecoding/soil/pkg/band/response"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustGeneralSettings() *response.GeneralSettings {
	result, e := c.GeneralSettings()
	errors.PanicOnError(e)

	return result
}
