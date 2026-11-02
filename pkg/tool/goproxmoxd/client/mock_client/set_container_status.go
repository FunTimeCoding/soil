package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"

func (c *Client) setContainerStatus(
	identifier int64,
	status string,
) (*client.TaskResult, error) {
	i, e := c.container(identifier)

	if e != nil {
		return nil, e
	}

	c.ContainerList[i].Status = &status

	return c.task(), nil
}
