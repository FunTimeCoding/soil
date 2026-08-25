package collector

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func newNode(registry *prometheus.Registry) (*Node, []*prometheus.GaugeVec) {
	result := &Node{
		status: gauge(
			registry,
			"proxmox_node_status",
			"Node status, one series per observed status with value 1",
			withStatus(nodeLabel()),
		),
		processor: gauge(
			registry,
			"proxmox_node_processor_ratio",
			"Node processor utilization between 0 and 1",
			nodeLabel(),
		),
		processorCount: gauge(
			registry,
			"proxmox_node_processor_count",
			"Number of processors available to the node",
			nodeLabel(),
		),
		memoryUsed: gauge(
			registry,
			"proxmox_node_memory_used_bytes",
			"Memory used by the node in bytes",
			nodeLabel(),
		),
		memoryTotal: gauge(
			registry,
			"proxmox_node_memory_total_bytes",
			"Memory available to the node in bytes",
			nodeLabel(),
		),
		diskUsed: gauge(
			registry,
			"proxmox_node_disk_used_bytes",
			"Root filesystem space used by the node in bytes",
			nodeLabel(),
		),
		diskTotal: gauge(
			registry,
			"proxmox_node_disk_total_bytes",
			"Root filesystem size of the node in bytes",
			nodeLabel(),
		),
		uptime: gauge(
			registry,
			"proxmox_node_uptime_seconds",
			"Node uptime in seconds",
			nodeLabel(),
		),
		version: gauge(
			registry,
			"proxmox_node_version_info",
			"Proxmox version of the node, always 1",
			withLabel(
				nodeLabel(),
				constant.ReleaseLabel,
				constant.RepositoryLabel,
				constant.VersionLabel,
			),
		),
		updatePending: gauge(
			registry,
			"proxmox_node_update_pending",
			"Number of package upgrades pending on the node",
			nodeLabel(),
		),
	}

	return result, []*prometheus.GaugeVec{
		result.status,
		result.processor,
		result.processorCount,
		result.memoryUsed,
		result.memoryTotal,
		result.diskUsed,
		result.diskTotal,
		result.uptime,
		result.version,
		result.updatePending,
	}
}
