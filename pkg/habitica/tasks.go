package habitica

import (
	"github.com/funtimecoding/soil/pkg/habitica/constant"
	"github.com/funtimecoding/soil/pkg/habitica/task"
	"net/url"
)

func (c *Client) Tasks(taskType string) ([]*task.Task, error) {
	v := url.Values{}

	if taskType != "" {
		v.Set(constant.TypeParameter, taskType)
	}

	var result []*task.Task
	e := c.basic.Get("/tasks/user", v, &result)

	return result, e
}
