package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func MachineDetail(v *proxmox.VirtualMachine) *server.MachineDetail {
	result := &server.MachineDetail{
		Identifier: int64(v.VMID),
		Name:       v.Name,
		Node:       &v.Node,
		Status:     &v.Status,
		Cpu:        new(v.CPU),
		Cpus:       new(v.CPUs),
		Mem:        new(int64(v.Mem)),
		MaxMem:     new(int64(v.MaxMem)),
		Disk:       new(int64(v.Disk)),
		MaxDisk:    new(int64(v.MaxDisk)),
		Uptime:     new(int64(v.Uptime)),
	}

	if v.Tags != "" {
		result.Tags = &v.Tags
	}

	if v.VirtualMachineConfig != nil {
		if v.VirtualMachineConfig.Description != "" {
			result.Description = &v.VirtualMachineConfig.Description
		}

		if v.VirtualMachineConfig.OSType != nil {
			result.OsType = v.VirtualMachineConfig.OSType
		}

		if v.VirtualMachineConfig.Sockets != nil {
			result.Sockets = v.VirtualMachineConfig.Sockets
		}

		if v.VirtualMachineConfig.Cores != nil {
			result.Cores = v.VirtualMachineConfig.Cores
		}
	}

	return result
}
