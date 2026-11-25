package mock_client

import (
	"context"
	"slices"
)

func (c *Client) MergeRequestApprovals(
	_ context.Context,
	project int64,
	identifier int64,
) ([]string, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	result := slices.Clone(c.approvals[[2]int64{project, identifier}])
	slices.Sort(result)

	return result, nil
}
