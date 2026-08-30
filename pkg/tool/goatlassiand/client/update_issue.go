package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) UpdateIssue(
	key string,
	summary string,
	description string,
	assignee string,
	reporter string,
	labels []string,
	fields map[string]any,
	noDiff bool,
) (string, int) {
	body := client.UpdateIssueJSONRequestBody{}

	if summary != "" {
		body.Summary = &summary
	}

	if description != "" {
		body.Description = &description
	}

	if assignee != "" {
		body.Assignee = &assignee
	}

	if reporter != "" {
		body.Reporter = &reporter
	}

	if len(labels) > 0 {
		body.Labels = &labels
	}

	if len(fields) > 0 {
		body.AdditionalFields = &fields
	}

	if noDiff {
		body.NoDiff = &noDiff
	}

	result, e := c.client.UpdateIssue(c.context, key, body)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
