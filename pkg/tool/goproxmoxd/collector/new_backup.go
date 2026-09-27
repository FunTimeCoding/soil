package collector

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func newBackup(
	registry *prometheus.Registry,
) (*Backup, []*prometheus.GaugeVec) {
	result := &Backup{
		missing: gauge(
			registry,
			"proxmox_guest_backup_missing",
			"Present when the guest is covered by no backup job",
			constant.BackupLabels,
		),
		missingCount: gauge(
			registry,
			"proxmox_backup_missing_count",
			"Number of guests covered by no backup job",
			constant.HypervisorLabels,
		),
	}

	return result, []*prometheus.GaugeVec{result.missing, result.missingCount}
}
