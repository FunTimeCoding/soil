package gitlab

import (
	"context"
	"fmt"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"strconv"
)

func (c *Client) ResolveProject(
	x context.Context,
	identifier string,
) (int64, error) {
	i, e := strconv.ParseInt(identifier, 10, 64)

	if e == nil {
		return i, nil
	}

	result, _, f := c.client.Projects.GetProject(
		identifier,
		nil,
		gitlab.WithContext(x),
	)

	if f != nil {
		return 0, fmt.Errorf("project %s: %w", identifier, f)
	}

	return result.ID, nil
}
