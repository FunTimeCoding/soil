package metric

import "github.com/prometheus/client_golang/prometheus"

type Backup struct {
	Missing      *prometheus.GaugeVec
	MissingCount *prometheus.GaugeVec
}
