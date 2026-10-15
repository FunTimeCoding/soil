package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateChoiceSet(
	name string,
	choices []string,
) *response.Response {
	result, e := c.client.CreateChoiceSet(
		c.context,
		client.CreateChoiceSetRequest{Name: name, Choices: choices},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
