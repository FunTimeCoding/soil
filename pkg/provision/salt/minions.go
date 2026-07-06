package salt

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response"

func (c *Client) Minions() ([]response.Minion, error) {
	return c.basic.ListMinions()
}
