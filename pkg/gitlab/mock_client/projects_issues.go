package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
)

func (c *Client) ProjectsIssues(_ context.Context) ([]*issue.Issue, error) {
	return nil, c.failure
}
