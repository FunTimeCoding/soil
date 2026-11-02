package mock_client

func (c *Client) SetMachineStatus(
	node string,
	identifier int,
	status string,
) {
	if nodeMachines, okay := c.machines[node]; okay {
		if vm, okay := nodeMachines[identifier]; okay {
			vm.Status = status
		}
	}
}
