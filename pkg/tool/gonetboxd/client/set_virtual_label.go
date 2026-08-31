package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) SetVirtualLabel(
	machine string,
	key string,
	value string,
) *response.Response {
	result, e := c.client.SetVirtualLabel(
		c.context,
		machine,
		key,
		client.SetVirtualLabelJSONRequestBody{Value: value},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
