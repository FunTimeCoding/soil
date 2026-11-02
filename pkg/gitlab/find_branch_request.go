package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) FindBranchRequest(
	x context.Context,
	project int64,
	branch string,
) (*merge_request.Request, bool, error) {
	result, e := c.BranchRequest(x, project, branch)

	if e != nil {
		if not_found.Is(e) {
			return nil, false, nil
		}

		return nil, false, e
	}

	return result, true, nil
}
