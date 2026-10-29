package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetTasks(
	x context.Context,
	taskType string,
) ([]*client.Task, error) {
	p := &client.GetTasksParams{}

	if taskType != "" {
		p.Type = new(client.GetTasksParamsType(taskType))
	}

	r, e := c.client.GetTasksWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get tasks", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
