package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) SearchUsers(
	x context.Context,
	query string,
) ([]*client.JiraUser, error) {
	r, e := c.client.SearchUsersWithResponse(
		x,
		&client.SearchUsersParams{Query: query},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("search users", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
