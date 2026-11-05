package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"slices"
)

func (c *Client) GetClaims(context.Context) ([]client.Claim, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if len(c.ClaimList) == 0 {
		return []client.Claim{}, nil
	}

	return slices.Clone(c.ClaimList), nil
}
