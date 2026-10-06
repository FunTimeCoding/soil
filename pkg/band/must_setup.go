package band

import (
	"github.com/funtimecoding/soil/pkg/band/response"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSetup() *response.Setup {
	result, e := c.Setup()
	errors.PanicOnError(e)

	return result
}
