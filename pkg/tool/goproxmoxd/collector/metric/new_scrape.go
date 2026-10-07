package metric

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func NewScrape(registry *prometheus.Registry) *Scrape {
	return &Scrape{
		Success: gauge(
			registry,
			"proxmox_scrape_success",
			"Whether the last poll of the hypervisor succeeded",
			constant.HypervisorLabels,
		),
		Duration: gauge(
			registry,
			"proxmox_scrape_duration_seconds",
			"Duration of the last poll of the hypervisor in seconds",
			constant.HypervisorLabels,
		),
	}
}
