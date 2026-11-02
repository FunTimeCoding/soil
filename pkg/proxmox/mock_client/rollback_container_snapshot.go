package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) RollbackContainerSnapshot(
	_ *proxmox.Container,
	_ string,
) (*proxmox.Task, error) {
	return &proxmox.Task{UPID: "mock:ct-rollback"}, nil
}
