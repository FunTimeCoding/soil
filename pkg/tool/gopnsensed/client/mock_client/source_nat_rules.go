package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) SourceNatRules(
	_ context.Context,
	_ string,
) ([]client.SourceNatRule, error) {
	return c.SourceNatList, c.Failure
}
