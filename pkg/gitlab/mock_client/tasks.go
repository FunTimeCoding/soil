package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/task"
)

func (c *Client) Tasks(_ context.Context) ([]*task.Task, error) {
	return nil, c.failure
}
