package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) SearchIssues(
	x context.Context,
	query string,
	limit *int,
) ([]*client.JiraIssue, error) {
	r, e := c.client.SearchIssuesWithResponse(
		x,
		&client.SearchIssuesParams{Query: query, Limit: limit},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("search issues", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
