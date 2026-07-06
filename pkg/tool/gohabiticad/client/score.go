package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Score(
	identifier string,
	direction string,
) string {
	result, e := c.client.ScoreTask(
		c.context,
		identifier,
		client.ScoreTaskParamsDirection(direction),
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
