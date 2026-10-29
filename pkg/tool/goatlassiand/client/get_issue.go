package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetIssue(
	x context.Context,
	key string,
	comments bool,
) (*client.JiraIssue, error) {
	params := &client.GetIssueParams{}

	if comments {
		params.Comments = &comments
	}

	r, e := c.client.GetIssueWithResponse(x, key, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("issue", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
