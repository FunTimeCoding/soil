package collector

import "github.com/prometheus/client_golang/prometheus"

type Backup struct {
	missing      *prometheus.GaugeVec
	missingCount *prometheus.GaugeVec
}
