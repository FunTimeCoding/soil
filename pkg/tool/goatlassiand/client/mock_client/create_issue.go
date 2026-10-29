package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) CreateIssue(
	_ context.Context,
	_ string,
	_ string,
	_ string,
	_ string,
	_ string,
	_ []string,
	_ map[string]any,
) (*client.JiraIssue, error) {
	return nil, c.Failure
}
