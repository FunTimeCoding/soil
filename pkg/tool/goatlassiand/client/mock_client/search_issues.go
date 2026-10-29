package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) SearchIssues(
	context.Context,
	string,
	*int,
) ([]*client.JiraIssue, error) {
	c.Searches++

	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Found, nil
}
