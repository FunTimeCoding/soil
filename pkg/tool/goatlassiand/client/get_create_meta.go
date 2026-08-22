package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) GetCreateMeta(
	project string,
	issueType string,
	expand string,
) string {
	params := &client.GetCreateMetaParams{
		Project:   project,
		IssueType: issueType,
	}

	if expand != "" {
		params.Expand = &expand
	}

	result, e := c.client.GetCreateMeta(c.context, params)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
