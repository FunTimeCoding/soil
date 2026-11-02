package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) StartMachine(_ *proxmox.VirtualMachine) (*proxmox.Task, error) {
	return &proxmox.Task{UPID: "mock:start"}, nil
}
