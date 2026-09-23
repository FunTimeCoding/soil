package worker_tester

import (
	"fmt"
	"strings"
)

func ScrapeSuccess(value string) string {
	return strings.Join(
		[]string{
			"# HELP proxmox_scrape_success Whether the last poll of the hypervisor succeeded",
			"# TYPE proxmox_scrape_success gauge",
			fmt.Sprintf("proxmox_scrape_success{hypervisor=\"pve\"} %s", value),
			"",
		},
		"\n",
	)
}
