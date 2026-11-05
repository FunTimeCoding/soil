package mock_client

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response/local_return"

func (c *Client) Highstate(
	_ string,
) (map[string]local_return.LocalReturn, error) {
	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.highstates++

	return map[string]local_return.LocalReturn{}, nil
}
