package worker

import "github.com/funtimecoding/soil/pkg/proxmox/constant"

func (w *Worker) pollInstance(hypervisor string) error {
	c, e := w.service.Client(hypervisor)

	if e != nil {
		return e
	}

	resources, e := c.ClusterResources()

	if e != nil {
		return e
	}

	w.collector.Clear(hypervisor)

	for _, r := range resources {
		switch r.Type {
		case constant.NodeType:
			w.collector.SetNode(hypervisor, r)
			w.pollNode(c, hypervisor, r.Node)
		case constant.MachineType, constant.ContainerType:
			w.collector.SetGuest(hypervisor, r)
		case constant.StorageType:
			w.collector.SetStorage(hypervisor, r)
		}
	}

	guests, e := c.GuestsNotInBackup()

	if e != nil {
		return e
	}

	w.collector.SetBackupMissing(hypervisor, guests)

	return nil
}
