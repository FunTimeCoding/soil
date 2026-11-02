package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) StopMachine(_ *proxmox.VirtualMachine) (*proxmox.Task, error) {
	return &proxmox.Task{UPID: "mock:stop"}, nil
}
