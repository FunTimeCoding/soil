package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"slices"
)

func New(registry *prometheus.Registry) *Collector {
	node, nodeVector := newNode(registry)
	guest, guestVector := newGuest(registry)
	storage, storageVector := newStorage(registry)
	backup, backupVector := newBackup(registry)

	return &Collector{
		node:    node,
		guest:   guest,
		storage: storage,
		backup:  backup,
		scrape:  newScrape(registry),
		clearable: slices.Concat(
			nodeVector,
			guestVector,
			storageVector,
			backupVector,
		),
	}
}
