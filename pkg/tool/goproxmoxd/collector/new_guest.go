package collector

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func newGuest(registry *prometheus.Registry) (*Guest, []*prometheus.GaugeVec) {
	result := &Guest{
		status: gauge(
			registry,
			"proxmox_guest_status",
			"Guest status, one series per observed status with value 1",
			withStatus(guestLabel()),
		),
		template: gauge(
			registry,
			"proxmox_guest_template",
			"Whether the guest is a template",
			guestLabel(),
		),
		tag: gauge(
			registry,
			"proxmox_guest_tag",
			"One series per tag assigned to the guest, always 1",
			withLabel(guestLabel(), constant.TagLabel),
		),
		processor: gauge(
			registry,
			"proxmox_guest_processor_ratio",
			"Guest processor utilization between 0 and 1",
			guestLabel(),
		),
		processorCount: gauge(
			registry,
			"proxmox_guest_processor_count",
			"Number of processors assigned to the guest",
			guestLabel(),
		),
		memoryUsed: gauge(
			registry,
			"proxmox_guest_memory_used_bytes",
			"Memory used by the guest in bytes",
			guestLabel(),
		),
		memoryTotal: gauge(
			registry,
			"proxmox_guest_memory_total_bytes",
			"Memory assigned to the guest in bytes",
			guestLabel(),
		),
		diskUsed: gauge(
			registry,
			"proxmox_guest_disk_used_bytes",
			"Root image space used by the guest in bytes",
			guestLabel(),
		),
		diskTotal: gauge(
			registry,
			"proxmox_guest_disk_total_bytes",
			"Root image size of the guest in bytes",
			guestLabel(),
		),
		uptime: gauge(
			registry,
			"proxmox_guest_uptime_seconds",
			"Guest uptime in seconds",
			guestLabel(),
		),
		networkReceive: gauge(
			registry,
			"proxmox_guest_network_receive_bytes",
			"Bytes received by the guest since it was started",
			guestLabel(),
		),
		networkTransmit: gauge(
			registry,
			"proxmox_guest_network_transmit_bytes",
			"Bytes sent by the guest since it was started",
			guestLabel(),
		),
		diskRead: gauge(
			registry,
			"proxmox_guest_disk_read_bytes",
			"Bytes read from block devices since the guest was started",
			guestLabel(),
		),
		diskWritten: gauge(
			registry,
			"proxmox_guest_disk_written_bytes",
			"Bytes written to block devices since the guest was started",
			guestLabel(),
		),
	}

	return result, []*prometheus.GaugeVec{
		result.status,
		result.template,
		result.tag,
		result.processor,
		result.processorCount,
		result.memoryUsed,
		result.memoryTotal,
		result.diskUsed,
		result.diskTotal,
		result.uptime,
		result.networkReceive,
		result.networkTransmit,
		result.diskRead,
		result.diskWritten,
	}
}
