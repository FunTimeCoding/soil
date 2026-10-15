package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateConfigurationContext(
	body client.CreateConfigurationContextRequest,
) *response.Response {
	result, e := c.client.CreateConfigurationContext(c.context, body)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
