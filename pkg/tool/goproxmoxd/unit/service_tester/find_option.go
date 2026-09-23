package service_tester

import "github.com/luthermonson/go-proxmox"

func FindOption(
	options []proxmox.VirtualMachineOption,
	name string,
) (any, bool) {
	for _, o := range options {
		if o.Name == name {
			return o.Value, true
		}
	}

	return nil, false
}
