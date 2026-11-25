package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
)

func (c *Client) MergeRequest(
	_ context.Context,
	project int64,
	identifier int64,
) (*merge_request_detail.Detail, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	for _, v := range c.details {
		if v.ProjectID == project && v.IID == identifier {
			return merge_request_detail.New(v), nil
		}
	}

	return nil, not_found.Format("merge request %d!%d", project, identifier)
}
