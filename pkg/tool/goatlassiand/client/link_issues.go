package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) LinkIssues(
	key string,
	target string,
	linkType string,
) *response.Response {
	body := client.LinkIssuesJSONRequestBody{TargetKey: target}

	if linkType != "" {
		body.LinkType = &linkType
	}

	result, e := c.client.LinkIssues(c.context, key, body)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
