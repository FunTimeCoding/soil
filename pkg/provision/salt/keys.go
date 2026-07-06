package salt

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response"

func (c *Client) Keys() (*response.KeysReturn, error) {
	return c.basic.ListKeys()
}
