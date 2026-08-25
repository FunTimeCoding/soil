package collector

import (
	"github.com/luthermonson/go-proxmox"
	"strconv"
)

func (c *Collector) SetBackupMissing(
	hypervisor string,
	guests []*proxmox.BackupGuestEntry,
) {
	for _, g := range guests {
		c.backup.missing.WithLabelValues(
			hypervisor,
			g.Type,
			strconv.Itoa(g.VMID),
			g.Name,
		).Set(1)
	}

	c.backup.missingCount.WithLabelValues(hypervisor).Set(float64(len(guests)))
}
