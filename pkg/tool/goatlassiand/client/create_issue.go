package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateIssue(
	x context.Context,
	project string,
	issueType string,
	summary string,
	description string,
	assignee string,
	labels []string,
	fields map[string]any,
) (*client.JiraIssue, error) {
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

	r, e := c.client.CreateIssueWithResponse(x, body)

	if e != nil {
		return nil, e
	}

	if r.JSON201 == nil {
		return nil, refusal.New("create issue", r.HTTPResponse, r.Body)
	}

	return r.JSON201, nil
}
