package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Metric struct {
	registry *prometheus.Registry
	upstream *prometheus.HistogramVec
	once     sync.Once
}
