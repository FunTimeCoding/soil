package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateTask(
	x context.Context,
	taskType string,
	text string,
	notes string,
) (*client.Task, error) {
	body := client.CreateTaskJSONRequestBody{
		Type: client.CreateTaskRequestType(taskType),
		Text: text,
	}

	if notes != "" {
		body.Notes = &notes
	}

	r, e := c.client.CreateTaskWithResponse(x, body)

	if e != nil {
		return nil, e
	}

	if r.JSON201 == nil {
		return nil, refusal.New("create task", r.HTTPResponse, r.Body)
	}

	return r.JSON201, nil
}
