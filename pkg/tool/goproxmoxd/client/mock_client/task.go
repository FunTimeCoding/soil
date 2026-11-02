package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) task() *client.TaskResult {
	c.tasks++

	return &client.TaskResult{TaskId: fmt.Sprintf("task-%d", c.tasks)}
}
