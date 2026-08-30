package raid

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Reports() (string, int) {
	result, e := c.client.GetReports(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
