package metric

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	pruned prometheus.Counter
}
