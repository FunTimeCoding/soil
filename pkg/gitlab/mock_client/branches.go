package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/branch"

func (c *Client) Branches(_ int64) ([]*branch.Branch, error) {
	return nil, nil
}
