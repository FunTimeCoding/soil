package salt

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response"

func (c *Client) Jobs() ([]response.Job, error) {
	return c.basic.ListJobs()
}
