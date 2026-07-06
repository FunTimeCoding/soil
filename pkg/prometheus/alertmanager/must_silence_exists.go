package alertmanager

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustSilenceExists(name string) bool {
	result, e := c.SilenceExists(name)
	errors.PanicOnError(e)

	return result
}
