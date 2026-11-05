package mock_client

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response/local_return"

func (c *Client) Local(
	_ string,
	_ string,
	_ []string,
) (map[string]local_return.LocalReturn, error) {
	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return map[string]local_return.LocalReturn{}, nil
}
