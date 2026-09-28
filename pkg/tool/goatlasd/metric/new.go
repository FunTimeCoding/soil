package metric

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func New(r *prometheus.Registry) *Metric {
	label := []string{constant.MetricSourceLabel}
	m := &Metric{
		placements: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "atlas_placements",
				Help: "Placements a source attributed in its last cycle",
			},
			label,
		),
		sightings: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "atlas_sightings",
				Help: "Sightings a source reported in its last cycle",
			},
			label,
		),
		sweptPlacements: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "atlas_swept_placements_total",
				Help: "Placements swept after passing retention",
			},
			label,
		),
		sweptSightings: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "atlas_swept_sightings_total",
				Help: "Sightings swept after passing retention",
			},
			label,
		),
		failures: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "atlas_collect_failures_total",
				Help: "Cycles a source failed to collect or save",
			},
			label,
		),
		seconds: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "atlas_collect_seconds",
				Help:    "Seconds a source takes to collect and save",
				Buckets: prometheus.ExponentialBuckets(0.1, 2, 10),
			},
			label,
		),
	}
	r.MustRegister(
		m.placements,
		m.sightings,
		m.sweptPlacements,
		m.sweptSightings,
		m.failures,
		m.seconds,
	)

	return m
}
