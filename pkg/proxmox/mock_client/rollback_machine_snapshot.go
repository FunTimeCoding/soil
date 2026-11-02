package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) RollbackMachineSnapshot(
	_ *proxmox.VirtualMachine,
	_ string,
) (*proxmox.Task, error) {
	return &proxmox.Task{UPID: "mock:rollback"}, nil
}
