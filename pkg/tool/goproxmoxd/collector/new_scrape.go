package collector

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func newScrape(registry *prometheus.Registry) *Scrape {
	return &Scrape{
		success: gauge(
			registry,
			"proxmox_scrape_success",
			"Whether the last poll of the hypervisor succeeded",
			constant.HypervisorLabels,
		),
		duration: gauge(
			registry,
			"proxmox_scrape_duration_seconds",
			"Duration of the last poll of the hypervisor in seconds",
			constant.HypervisorLabels,
		),
	}
}
