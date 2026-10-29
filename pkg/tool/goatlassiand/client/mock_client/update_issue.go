package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) UpdateIssue(
	_ context.Context,
	_ string,
	_ string,
	_ string,
	_ string,
	_ string,
	_ []string,
	_ map[string]any,
	_ bool,
) (*client.IssueUpdateResult, error) {
	return nil, c.Failure
}
