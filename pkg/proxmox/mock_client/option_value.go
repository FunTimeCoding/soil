package mock_client

import "github.com/luthermonson/go-proxmox"

func optionValue(
	options []proxmox.VirtualMachineOption,
	name string,
) string {
	for _, o := range options {
		if o.Name == name {
			if s, okay := o.Value.(string); okay {
				return s
			}
		}
	}

	return ""
}
