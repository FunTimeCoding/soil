package openbao

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustPut(
	path string,
	values map[string]string,
) {
	errors.PanicOnError(c.Put(path, values))
}
