package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
)

func (c *Client) CreateTask(
	_ context.Context,
	taskType string,
	text string,
	notes string,
) (*client.Task, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result := &client.Task{Type: taskType, Text: text}

	if notes != "" {
		result.Notes = &notes
	}

	c.CreatedList = append(c.CreatedList, result)

	return result, nil
}
