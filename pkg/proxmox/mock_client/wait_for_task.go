package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) WaitForTask(
	_ *proxmox.Task,
	_ int,
) error {
	return nil
}
