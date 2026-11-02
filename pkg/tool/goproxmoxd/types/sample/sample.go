package sample

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/node_sample"
	"github.com/luthermonson/go-proxmox"
	"time"
)

type Sample struct {
	Hypervisor string
	Resources  proxmox.ClusterResources
	Nodes      []*node_sample.NodeSample
	Unbacked   []*proxmox.BackupGuestEntry
	Duration   time.Duration
	Failure    error
}
