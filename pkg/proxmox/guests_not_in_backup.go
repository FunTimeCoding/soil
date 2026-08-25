package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) GuestsNotInBackup() ([]*proxmox.BackupGuestEntry, error) {
	cluster, e := c.Cluster()

	if e != nil {
		return nil, e
	}

	return cluster.GuestsNotInBackup(c.context)
}
