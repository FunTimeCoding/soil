package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListPages(
	space string,
	status string,
) *response.Response {
	params := &client.ListPagesParams{Space: space}

	if status != "" {
		params.Status = &status
	}

	result, e := c.client.ListPages(c.context, params)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
