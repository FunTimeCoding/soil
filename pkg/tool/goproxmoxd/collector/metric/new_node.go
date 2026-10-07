package metric

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func NewNode(registry *prometheus.Registry) (*Node, []*prometheus.GaugeVec) {
	result := &Node{
		Status: gauge(
			registry,
			"proxmox_node_status",
			"Node status, one series per observed status with value 1",
			withStatus(constant.NodeLabels),
		),
		Processor: gauge(
			registry,
			"proxmox_node_processor_ratio",
			"Node processor utilization between 0 and 1",
			constant.NodeLabels,
		),
		ProcessorCount: gauge(
			registry,
			"proxmox_node_processor_count",
			"Number of processors available to the node",
			constant.NodeLabels,
		),
		MemoryUsed: gauge(
			registry,
			"proxmox_node_memory_used_bytes",
			"Memory used by the node in bytes",
			constant.NodeLabels,
		),
		MemoryTotal: gauge(
			registry,
			"proxmox_node_memory_total_bytes",
			"Memory available to the node in bytes",
			constant.NodeLabels,
		),
		DiskUsed: gauge(
			registry,
			"proxmox_node_disk_used_bytes",
			"Root filesystem space used by the node in bytes",
			constant.NodeLabels,
		),
		DiskTotal: gauge(
			registry,
			"proxmox_node_disk_total_bytes",
			"Root filesystem size of the node in bytes",
			constant.NodeLabels,
		),
		Uptime: gauge(
			registry,
			"proxmox_node_uptime_seconds",
			"Node uptime in seconds",
			constant.NodeLabels,
		),
		Version: gauge(
			registry,
			"proxmox_node_version_info",
			"Proxmox version of the node, always 1",
			WithLabel(
				constant.NodeLabels,
				constant.ReleaseLabel,
				constant.RepositoryLabel,
				constant.VersionLabel,
			),
		),
		UpdatePending: gauge(
			registry,
			"proxmox_node_update_pending",
			"Number of package upgrades pending on the node",
			constant.NodeLabels,
		),
	}

	return result, []*prometheus.GaugeVec{
		result.Status,
		result.Processor,
		result.ProcessorCount,
		result.MemoryUsed,
		result.MemoryTotal,
		result.DiskUsed,
		result.DiskTotal,
		result.Uptime,
		result.Version,
		result.UpdatePending,
	}
}
