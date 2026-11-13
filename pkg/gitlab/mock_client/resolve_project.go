package mock_client

import (
	"context"
	"strconv"
)

func (c *Client) ResolveProject(
	_ context.Context,
	identifier string,
) (int64, error) {
	if c.failure != nil {
		return 0, c.failure
	}

	if number, e := strconv.ParseInt(identifier, 10, 64); e == nil {
		return number, nil
	}

	for _, p := range c.projects {
		if p.CombinedName() == identifier {
			return p.Identifier, nil
		}
	}

	return 0, nil
}
