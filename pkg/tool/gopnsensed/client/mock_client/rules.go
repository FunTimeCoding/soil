package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) Rules(
	_ context.Context,
	_ string,
) ([]client.Rule, error) {
	return c.RuleList, c.Failure
}
