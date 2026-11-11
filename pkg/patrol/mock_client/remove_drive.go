package mock_client

import "github.com/funtimecoding/soil/pkg/smartctl/response"

func (c *Client) RemoveDrive(name string) {
	var kept []*response.Report

	for _, r := range c.drives {
		if r.Device.Name != name {
			kept = append(kept, r)
		}
	}

	c.drives = kept
}
