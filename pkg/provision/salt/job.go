package salt

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response"

func (c *Client) Job(identifier string) (*response.Job, error) {
	return c.basic.LookupJob(identifier)
}
