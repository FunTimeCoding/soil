package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/task"

func (c *Client) Tasks() ([]*task.Task, error) {
	return nil, c.failure
}
