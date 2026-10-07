package collector

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/collector/metric"
	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	node      *metric.Node
	guest     *metric.Guest
	storage   *metric.Storage
	backup    *metric.Backup
	scrape    *metric.Scrape
	clearable []*prometheus.GaugeVec
}
