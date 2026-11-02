package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) AddStorage(
	node string,
	name string,
	storageType string,
	content string,
) {
	c.storages[node] = append(
		c.storages[node],
		&proxmox.Storage{
			Name:    name,
			Type:    storageType,
			Content: content,
			Enabled: 1,
			Active:  1,
		},
	)
}
