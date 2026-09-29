package prometheus

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustAllLabels() []string {
	result, e := c.AllLabels()
	errors.PanicOnError(e)

	return result
}
