package mock_client

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response"

func (c *Client) Minions() ([]response.Minion, error) {
	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return []response.Minion{}, nil
}
