package mock_client

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response"

func (c *Client) Jobs() ([]response.Job, error) {
	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return []response.Job{}, nil
}
