package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/task"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Tasks() ([]*task.Task, error) {
	result, _, e := c.client.Todos.ListTodos(
		&gitlab.ListTodosOptions{
			State:       new(constant.PendingState),
			ListOptions: constant.DefaultListOptions,
		},
		nil,
	)

	if e != nil {
		return nil, e
	}

	return task.NewSlice(result), nil
}
