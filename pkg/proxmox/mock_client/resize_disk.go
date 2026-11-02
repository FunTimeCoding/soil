package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) ResizeDisk(
	_ *proxmox.VirtualMachine,
	_ string,
	_ string,
) (*proxmox.Task, error) {
	return &proxmox.Task{UPID: "mock:resize"}, nil
}
