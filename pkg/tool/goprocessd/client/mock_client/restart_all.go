package mock_client

import "context"

func (c *Client) RestartAll(context.Context) (int, error) {
	if c.Failure != nil {
		return 0, c.Failure
	}

	for i := range c.ProcessList {
		c.ProcessList[i].Running = true
	}

	return len(c.ProcessList), nil
}
