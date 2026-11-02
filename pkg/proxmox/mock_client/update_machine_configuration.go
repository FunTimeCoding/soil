package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) UpdateMachineConfiguration(
	_ *proxmox.VirtualMachine,
	_ ...proxmox.VirtualMachineOption,
) error {
	return nil
}
