package runner

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func newMetrics(
	r *prometheus.Registry,
	held func() *lease.Lease,
) *metrics {
	m := &metrics{
		runsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "terraform_run_total",
				Help: "Total number of terraform applies by outcome.",
			},
			[]string{constant.RunnerStatus},
		),
		applyDuration: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name:    "terraform_apply_duration_seconds",
				Help:    "Duration of a single terraform apply.",
				Buckets: prometheus.DefBuckets,
			},
		),
		lastSuccess: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "terraform_last_success_timestamp_seconds",
				Help: "Unix timestamp of the last successful terraform apply.",
			},
		),
	}
	r.MustRegister(
		m.runsTotal,
		m.applyDuration,
		m.lastSuccess,
		prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: "terraform_state_locked",
				Help: "Whether the terraform state lock is currently held.",
			},
			func() float64 {
				if v := held(); v != nil && v.Held() {
					return 1
				}

				return 0
			},
		),
		prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: "terraform_state_lock_age_seconds",
				Help: "How long the terraform state lock has been held.",
			},
			func() float64 {
				v := held()

				if v == nil {
					return 0
				}

				return v.Age().Seconds()
			},
		),
	)

	return m
}
