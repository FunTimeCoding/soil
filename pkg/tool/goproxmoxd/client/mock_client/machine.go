package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) machine(identifier int64) (int, error) {
	if c.Failure != nil {
		return 0, c.Failure
	}

	for i, m := range c.MachineList {
		if m.Identifier == identifier {
			return i, nil
		}
	}

	return 0, not_found.Format("machine %d", identifier)
}
