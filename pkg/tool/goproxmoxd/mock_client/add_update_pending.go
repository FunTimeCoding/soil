package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) AddUpdatePending(
	node string,
	name string,
	oldVersion string,
	version string,
) {
	c.updatesPending[node] = append(
		c.updatesPending[node],
		&proxmox.APTUpdate{
			Package:    name,
			OldVersion: oldVersion,
			Version:    version,
		},
	)
}
