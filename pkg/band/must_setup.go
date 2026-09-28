package band

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustSetup() *SetupResponse {
	result, e := c.Setup()
	errors.PanicOnError(e)

	return result
}
