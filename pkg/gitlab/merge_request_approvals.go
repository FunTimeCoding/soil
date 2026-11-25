package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"slices"
)

func (c *Client) MergeRequestApprovals(
	x context.Context,
	project int64,
	identifier int64,
) ([]string, error) {
	v, _, e := c.client.MergeRequestApprovals.GetConfiguration(
		project,
		identifier,
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	result := make([]string, 0, len(v.ApprovedBy))

	for _, a := range v.ApprovedBy {
		if a.User != nil {
			result = append(result, a.User.Username)
		}
	}

	slices.Sort(result)

	return result, nil
}
