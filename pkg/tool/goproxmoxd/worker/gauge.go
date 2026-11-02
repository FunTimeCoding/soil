package worker

import (
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/sample"
)

func (w *Worker) gauge(v *sample.Sample) {
	for _, r := range v.Resources {
		switch r.Type {
		case proxmoxConstant.NodeType:
			w.collector.SetNode(v.Hypervisor, r)
		case proxmoxConstant.MachineType, proxmoxConstant.ContainerType:
			w.collector.SetGuest(v.Hypervisor, r)
		case proxmoxConstant.StorageType:
			w.collector.SetStorage(v.Hypervisor, r)
		}
	}

	for _, n := range v.Nodes {
		if n.Version != nil {
			w.collector.SetVersion(v.Hypervisor, n.Name, n.Version)
		}

		if n.UpdatesFailure == nil {
			w.collector.SetUpdatePending(v.Hypervisor, n.Name, len(n.Updates))
		}
	}

	w.collector.SetBackupMissing(v.Hypervisor, v.Unbacked)
}
