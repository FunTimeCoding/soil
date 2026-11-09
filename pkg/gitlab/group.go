package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Group(
	x context.Context,
	s string,
) (*gitlab.Group, error) {
	result, _, e := c.client.Groups.SearchGroup(s, gitlab.WithContext(x))

	if e != nil {
		return nil, wrapError(e)
	}

	if len(result) == 0 {
		return nil, not_found.New("group", s)
	}

	if len(result) > 1 {
		return nil, ambiguous.Format("expected 1 group, got %d", len(result))
	}

	return result[0], nil
}
