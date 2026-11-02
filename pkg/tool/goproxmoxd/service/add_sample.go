package service

import (
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/guest"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/node"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/storage"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/sample"
)

func addSample(
	f *floor.Floor,
	v *sample.Sample,
) {
	for _, n := range v.Nodes {
		result := node.Node{Hypervisor: v.Hypervisor, Name: n.Name}

		if n.Version != nil {
			result.Version = n.Version.Release
		}

		if n.UpdatesFailure == nil {
			result.UpdatesPending = len(n.Updates)
		}

		f.Nodes = append(f.Nodes, result)
	}

	missing := make(map[uint64]bool)

	for _, g := range v.Unbacked {
		missing[uint64(g.VMID)] = true
	}

	for _, r := range v.Resources {
		switch r.Type {
		case proxmoxConstant.MachineType, proxmoxConstant.ContainerType:
			if r.Template != 0 {
				continue
			}

			f.Guests = append(
				f.Guests,
				guest.Guest{
					Hypervisor:  v.Hypervisor,
					Node:        r.Node,
					Kind:        r.Type,
					Identifier:  r.VMID,
					Name:        r.Name,
					Status:      r.Status,
					Processor:   r.CPU,
					Memory:      r.Mem,
					MemoryTotal: r.MaxMem,
					Unbacked:    missing[r.VMID],
				},
			)
		case proxmoxConstant.StorageType:
			f.Storages = append(
				f.Storages,
				storage.Storage{
					Hypervisor: v.Hypervisor,
					Name:       r.Storage,
					Used:       r.Disk,
					Total:      r.MaxDisk,
				},
			)
		}
	}
}
