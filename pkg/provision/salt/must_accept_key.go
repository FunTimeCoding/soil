package salt

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustAcceptKey(minion string) []string {
	result, e := c.AcceptKey(minion)
	errors.PanicOnError(e)

	return result
}
