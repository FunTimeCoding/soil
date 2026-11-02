package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) AddGuestNotInBackup(
	guestType string,
	identifier int,
	name string,
) {
	c.notInBackup = append(
		c.notInBackup,
		&proxmox.BackupGuestEntry{
			VMID: identifier,
			Type: guestType,
			Name: name,
		},
	)
}
