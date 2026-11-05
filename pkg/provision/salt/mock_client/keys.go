package mock_client

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response"

func (c *Client) Keys() (*response.KeysReturn, error) {
	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return &c.keys, nil
}
