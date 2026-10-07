package worker

import (
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/guest"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/storage"
)

func (w *Worker) pollInstance(
	hypervisor string,
	f *floor.Floor,
) error {
	c, e := w.service.Client(hypervisor)

	if e != nil {
		return e
	}

	resources, e := c.ClusterResources()

	if e != nil {
		return e
	}

	w.collector.Clear(hypervisor)
	guests := []guest.Guest{}

	for _, r := range resources {
		switch r.Type {
		case constant.NodeType:
			w.collector.SetNode(hypervisor, r)
			f.Nodes = append(f.Nodes, *w.pollNode(c, hypervisor, r.Node))
		case constant.MachineType, constant.ContainerType:
			w.collector.SetGuest(hypervisor, r)

			if r.Template != 0 {
				continue
			}

			guests = append(
				guests,
				guest.Guest{
					Hypervisor:  hypervisor,
					Node:        r.Node,
					Kind:        r.Type,
					Identifier:  r.VMID,
					Name:        r.Name,
					Status:      r.Status,
					Processor:   r.CPU,
					Memory:      r.Mem,
					MemoryTotal: r.MaxMem,
				},
			)
		case constant.StorageType:
			w.collector.SetStorage(hypervisor, r)
			f.Storages = append(
				f.Storages,
				storage.Storage{
					Hypervisor: hypervisor,
					Name:       r.Storage,
					Used:       r.Disk,
					Total:      r.MaxDisk,
				},
			)
		}
	}

	unbacked, e := c.GuestsNotInBackup()

	if e != nil {
		return e
	}

	w.collector.SetBackupMissing(hypervisor, unbacked)
	missing := make(map[uint64]bool)

	for _, g := range unbacked {
		missing[uint64(g.VMID)] = true
	}

	for index := range guests {
		guests[index].Unbacked = missing[guests[index].Identifier]
	}

	f.Guests = append(f.Guests, guests...)

	return nil
}
