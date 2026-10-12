package openbao

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustRead(path string) map[string]string {
	result, e := c.Read(path)
	errors.PanicOnError(e)

	return result
}
