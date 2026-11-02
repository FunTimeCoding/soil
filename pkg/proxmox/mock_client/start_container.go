package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) StartContainer(_ *proxmox.Container) (*proxmox.Task, error) {
	return &proxmox.Task{UPID: "mock:ct-start"}, nil
}
