package raid

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Reports() string {
	result, e := c.client.GetReports(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
