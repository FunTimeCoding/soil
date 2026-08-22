package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateIssue(
	project string,
	issueType string,
	summary string,
	description string,
	assignee string,
	labels []string,
	fields map[string]any,
) string {
	body := client.CreateIssueJSONRequestBody{
		Project:   project,
		IssueType: issueType,
		Summary:   summary,
	}

	if description != "" {
		body.Description = &description
	}

	if assignee != "" {
		body.Assignee = &assignee
	}

	if len(labels) > 0 {
		body.Labels = &labels
	}

	if len(fields) > 0 {
		body.AdditionalFields = &fields
	}

	result, e := c.client.CreateIssue(c.context, body)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
