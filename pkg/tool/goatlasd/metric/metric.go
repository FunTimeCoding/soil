package metric

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	placements      *prometheus.GaugeVec
	sightings       *prometheus.GaugeVec
	sweptPlacements *prometheus.CounterVec
	sweptSightings  *prometheus.CounterVec
	failures        *prometheus.CounterVec
	seconds         *prometheus.HistogramVec
}
