package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) UpdateConfigurationContext(
	name string,
	body client.UpdateConfigurationContextRequest,
) *response.Response {
	result, e := c.client.UpdateConfigurationContext(c.context, name, body)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
