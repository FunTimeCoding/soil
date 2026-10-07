package metric

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/lease"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/lock_detail"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	r *prometheus.Registry,
	held func() *lease.Lease,
) *Metric {
	m := &Metric{
		RunsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "terraform_run_total",
				Help: "Total number of terraform applies by outcome.",
			},
			[]string{constant.RunnerStatus},
		),
		ApplyDuration: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name:    "terraform_apply_duration_seconds",
				Help:    "Duration of a single terraform apply.",
				Buckets: prometheus.DefBuckets,
			},
		),
		LastSuccess: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "terraform_last_success_timestamp_seconds",
				Help: "Unix timestamp of the last successful terraform apply.",
			},
		),
	}
	r.MustRegister(
		m.RunsTotal,
		m.ApplyDuration,
		m.LastSuccess,
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
				d := lock_detail.New(held())

				if d == nil {
					return 0
				}

				return time.Since(d.Created).Seconds()
			},
		),
	)

	return m
}
