package gitlab

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustQuery(
	q string,
	response any,
) {
	errors.PanicOnError(c.Query(q, response))
}
