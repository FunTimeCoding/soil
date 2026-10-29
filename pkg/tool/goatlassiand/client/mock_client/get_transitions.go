package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) GetTransitions(
	_ context.Context,
	_ string,
) ([]*client.JiraTransition, error) {
	return nil, c.Failure
}
