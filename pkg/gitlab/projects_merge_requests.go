package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) ProjectsMergeRequests(
	x context.Context,
) ([]*merge_request.Request, error) {
	var result []*merge_request.Request

	for _, identifier := range c.projects {
		requests, e := c.ProjectMergeRequests(x, identifier, "")

		if e != nil {
			return nil, e
		}

		for _, r := range requests {
			if r.Done() {
				continue
			}

			result = append(result, r)
		}
	}

	return result, nil
}
