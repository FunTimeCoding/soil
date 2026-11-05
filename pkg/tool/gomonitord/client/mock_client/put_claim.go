package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
)

func (c *Client) PutClaim(
	_ context.Context,
	item string,
	owner string,
) (*client.Claim, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if item == "" || owner == "" {
		return nil, validation.New("%s", constant.ClaimRequiredText)
	}

	for _, l := range c.ClaimList {
		if l.Item == item && l.Owner == owner {
			return &l, nil
		}
	}

	result := client.Claim{Item: item, Owner: owner}
	c.ClaimList = append(c.ClaimList, result)

	return &result, nil
}
