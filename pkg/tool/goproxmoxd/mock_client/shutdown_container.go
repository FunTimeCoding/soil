package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) ShutdownContainer(_ *proxmox.Container) (*proxmox.Task, error) {
	return &proxmox.Task{UPID: "mock:ct-shutdown"}, nil
}
