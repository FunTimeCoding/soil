package metric

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func NewGuest(registry *prometheus.Registry) (*Guest, []*prometheus.GaugeVec) {
	result := &Guest{
		Status: gauge(
			registry,
			"proxmox_guest_status",
			"Guest status, one series per observed status with value 1",
			withStatus(constant.GuestLabels),
		),
		Template: gauge(
			registry,
			"proxmox_guest_template",
			"Whether the guest is a template",
			constant.GuestLabels,
		),
		Tag: gauge(
			registry,
			"proxmox_guest_tag",
			"One series per tag assigned to the guest, always 1",
			WithLabel(constant.GuestLabels, constant.TagLabel),
		),
		Processor: gauge(
			registry,
			"proxmox_guest_processor_ratio",
			"Guest processor utilization between 0 and 1",
			constant.GuestLabels,
		),
		ProcessorCount: gauge(
			registry,
			"proxmox_guest_processor_count",
			"Number of processors assigned to the guest",
			constant.GuestLabels,
		),
		MemoryUsed: gauge(
			registry,
			"proxmox_guest_memory_used_bytes",
			"Memory used by the guest in bytes",
			constant.GuestLabels,
		),
		MemoryTotal: gauge(
			registry,
			"proxmox_guest_memory_total_bytes",
			"Memory assigned to the guest in bytes",
			constant.GuestLabels,
		),
		DiskUsed: gauge(
			registry,
			"proxmox_guest_disk_used_bytes",
			"Root image space used by the guest in bytes",
			constant.GuestLabels,
		),
		DiskTotal: gauge(
			registry,
			"proxmox_guest_disk_total_bytes",
			"Root image size of the guest in bytes",
			constant.GuestLabels,
		),
		Uptime: gauge(
			registry,
			"proxmox_guest_uptime_seconds",
			"Guest uptime in seconds",
			constant.GuestLabels,
		),
		NetworkReceive: gauge(
			registry,
			"proxmox_guest_network_receive_bytes",
			"Bytes received by the guest since it was started",
			constant.GuestLabels,
		),
		NetworkTransmit: gauge(
			registry,
			"proxmox_guest_network_transmit_bytes",
			"Bytes sent by the guest since it was started",
			constant.GuestLabels,
		),
		DiskRead: gauge(
			registry,
			"proxmox_guest_disk_read_bytes",
			"Bytes read from block devices since the guest was started",
			constant.GuestLabels,
		),
		DiskWritten: gauge(
			registry,
			"proxmox_guest_disk_written_bytes",
			"Bytes written to block devices since the guest was started",
			constant.GuestLabels,
		),
	}

	return result, []*prometheus.GaugeVec{
		result.Status,
		result.Template,
		result.Tag,
		result.Processor,
		result.ProcessorCount,
		result.MemoryUsed,
		result.MemoryTotal,
		result.DiskUsed,
		result.DiskTotal,
		result.Uptime,
		result.NetworkReceive,
		result.NetworkTransmit,
		result.DiskRead,
		result.DiskWritten,
	}
}
