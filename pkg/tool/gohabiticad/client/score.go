package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Score(
	identifier string,
	direction string,
) *response.Response {
	result, e := c.client.ScoreTask(
		c.context,
		identifier,
		client.ScoreTaskParamsDirection(direction),
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
