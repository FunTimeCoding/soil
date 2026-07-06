package alertmanager

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustSimpleSilence(alert string) string {
	result, e := c.SimpleSilence(alert)
	errors.PanicOnError(e)

	return result
}
