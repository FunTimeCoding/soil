package mock_client

import "github.com/funtimecoding/soil/pkg/smartctl/response"

func (c *Client) AddDrive(r ...*response.Report) {
	c.drives = append(c.drives, r...)
}
