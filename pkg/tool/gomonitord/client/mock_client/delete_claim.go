package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"slices"
)

func (c *Client) DeleteClaim(
	_ context.Context,
	item string,
	owner string,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	if item == "" || owner == "" {
		return validation.New("%s", constant.ClaimRequiredText)
	}

	c.ClaimList = slices.DeleteFunc(
		c.ClaimList,
		func(l client.Claim) bool { return l.Item == item && l.Owner == owner },
	)

	return nil
}
