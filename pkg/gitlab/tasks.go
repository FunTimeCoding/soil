package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/task"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Tasks(x context.Context) ([]*task.Task, error) {
	result, _, e := c.client.Todos.ListTodos(
		&gitlab.ListTodosOptions{
			State:       new(constant.PendingState),
			ListOptions: constant.DefaultListOptions,
		},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return task.NewSlice(result), nil
}
