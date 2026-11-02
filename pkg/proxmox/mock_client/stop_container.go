package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) StopContainer(_ *proxmox.Container) (*proxmox.Task, error) {
	return &proxmox.Task{UPID: "mock:ct-stop"}, nil
}
