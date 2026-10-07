package collector

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/collector/metric"
	"github.com/prometheus/client_golang/prometheus"
	"slices"
)

func New(registry *prometheus.Registry) *Collector {
	node, nodeVector := metric.NewNode(registry)
	guest, guestVector := metric.NewGuest(registry)
	storage, storageVector := metric.NewStorage(registry)
	backup, backupVector := metric.NewBackup(registry)

	return &Collector{
		node:    node,
		guest:   guest,
		storage: storage,
		backup:  backup,
		scrape:  metric.NewScrape(registry),
		clearable: slices.Concat(
			nodeVector,
			guestVector,
			storageVector,
			backupVector,
		),
	}
}
