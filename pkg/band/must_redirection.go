package band

import (
	"github.com/funtimecoding/soil/pkg/band/response"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustRedirection() *response.Redirection {
	result, e := c.Redirection()
	errors.PanicOnError(e)

	return result
}
