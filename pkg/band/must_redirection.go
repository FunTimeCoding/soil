package band

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustRedirection() *RedirectionResponse {
	result, e := c.Redirection()
	errors.PanicOnError(e)

	return result
}
