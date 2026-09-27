package collector

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func newStorage(
	registry *prometheus.Registry,
) (*Storage, []*prometheus.GaugeVec) {
	result := &Storage{
		status: gauge(
			registry,
			"proxmox_storage_status",
			"Storage status, one series per observed status with value 1",
			withStatus(constant.StorageLabels),
		),
		used: gauge(
			registry,
			"proxmox_storage_used_bytes",
			"Storage space used in bytes",
			constant.StorageLabels,
		),
		total: gauge(
			registry,
			"proxmox_storage_total_bytes",
			"Storage size in bytes",
			constant.StorageLabels,
		),
		shared: gauge(
			registry,
			"proxmox_storage_shared",
			"Whether the storage is shared among cluster nodes",
			constant.StorageLabels,
		),
	}

	return result, []*prometheus.GaugeVec{
		result.status,
		result.used,
		result.total,
		result.shared,
	}
}
