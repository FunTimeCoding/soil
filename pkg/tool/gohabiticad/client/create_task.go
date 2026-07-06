package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateTask(
	taskType string,
	text string,
	notes string,
) string {
	body := client.CreateTaskJSONRequestBody{
		Type: client.CreateTaskRequestType(taskType),
		Text: text,
	}

	if notes != "" {
		body.Notes = &notes
	}

	result, e := c.client.CreateTask(c.context, body)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
