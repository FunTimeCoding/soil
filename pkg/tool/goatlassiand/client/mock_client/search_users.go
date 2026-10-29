package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) SearchUsers(
	_ context.Context,
	_ string,
) ([]*client.JiraUser, error) {
	return nil, c.Failure
}
