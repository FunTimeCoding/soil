package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) UpdateIssue(
	x context.Context,
	key string,
	summary string,
	description string,
	assignee string,
	reporter string,
	labels []string,
	fields map[string]any,
	noDiff bool,
) (*client.IssueUpdateResult, error) {
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

	r, e := c.client.UpdateIssueWithResponse(x, key, body)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("update issue", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
