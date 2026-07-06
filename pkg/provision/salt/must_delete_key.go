package salt

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteKey(minion string) []string {
	result, e := c.DeleteKey(minion)
	errors.PanicOnError(e)

	return result
}
