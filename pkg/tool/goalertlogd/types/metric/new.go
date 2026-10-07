package metric

import "github.com/prometheus/client_golang/prometheus"

func New(r *prometheus.Registry) *Metric {
	m := &Metric{
		AlertsTotal: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "alertlog_alerts_total",
				Help: "Total number of alerts recorded.",
			},
		),
		AlertsFiring: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "alertlog_alerts_firing",
				Help: "Number of currently firing alerts.",
			},
		),
		RecordsTotal: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "alertlog_records_total",
				Help: "Total number of stored alert records.",
			},
		),
		PollDuration: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name:    "alertlog_poll_duration_seconds",
				Help:    "Duration of a single poll cycle.",
				Buckets: prometheus.DefBuckets,
			},
		),
		LastPollTime: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "alertlog_last_poll_timestamp_seconds",
				Help: "Unix timestamp of the last successful poll.",
			},
		),
	}
	r.MustRegister(
		m.AlertsTotal,
		m.AlertsFiring,
		m.RecordsTotal,
		m.PollDuration,
		m.LastPollTime,
	)

	return m
}
