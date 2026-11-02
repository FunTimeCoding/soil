package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) GuestsNotInBackup() ([]*proxmox.BackupGuestEntry, error) {
	return c.notInBackup, nil
}
