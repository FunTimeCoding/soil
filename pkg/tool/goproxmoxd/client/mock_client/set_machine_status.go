package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"

func (c *Client) setMachineStatus(
	identifier int64,
	status string,
) (*client.TaskResult, error) {
	i, e := c.machine(identifier)

	if e != nil {
		return nil, e
	}

	c.MachineList[i].Status = &status

	return c.task(), nil
}
