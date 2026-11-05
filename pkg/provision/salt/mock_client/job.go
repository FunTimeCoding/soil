package mock_client

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response"

func (c *Client) Job(_ string) (*response.Job, error) {
	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	var result response.Job

	return &result, nil
}
