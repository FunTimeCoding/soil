package collector

import "github.com/prometheus/client_golang/prometheus"

type Collector struct {
	node      *Node
	guest     *Guest
	storage   *Storage
	backup    *Backup
	scrape    *Scrape
	clearable []*prometheus.GaugeVec
}
