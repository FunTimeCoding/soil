package metric

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func NewStorage(
	registry *prometheus.Registry,
) (*Storage, []*prometheus.GaugeVec) {
	result := &Storage{
		Status: gauge(
			registry,
			"proxmox_storage_status",
			"Storage status, one series per observed status with value 1",
			withStatus(constant.StorageLabels),
		),
		Used: gauge(
			registry,
			"proxmox_storage_used_bytes",
			"Storage space used in bytes",
			constant.StorageLabels,
		),
		Total: gauge(
			registry,
			"proxmox_storage_total_bytes",
			"Storage size in bytes",
			constant.StorageLabels,
		),
		Shared: gauge(
			registry,
			"proxmox_storage_shared",
			"Whether the storage is shared among cluster nodes",
			constant.StorageLabels,
		),
	}

	return result, []*prometheus.GaugeVec{
		result.Status,
		result.Used,
		result.Total,
		result.Shared,
	}
}
