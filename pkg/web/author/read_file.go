package author

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) ReadFile(path string) []byte {
	result, e := c.client.Read(path)
	errors.PanicOnError(e)

	return result
}
