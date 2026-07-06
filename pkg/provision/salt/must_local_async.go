package salt

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustLocalAsync(
	target string,
	function string,
	a []string,
) string {
	result, e := c.LocalAsync(target, function, a)
	errors.PanicOnError(e)

	return result
}
