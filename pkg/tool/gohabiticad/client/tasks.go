package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Tasks(taskType string) string {
	var p *client.GetTasksParams

	if taskType != "" {
		p = &client.GetTasksParams{
			Type: new(client.GetTasksParamsType(taskType)),
		}
	}

	result, e := c.client.GetTasks(c.context, p)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
