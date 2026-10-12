package openbao

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustList(path string) []string {
	result, e := c.List(path)
	errors.PanicOnError(e)

	return result
}
