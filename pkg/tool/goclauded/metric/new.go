package metric

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func New(r *prometheus.Registry) *Metric {
	m := &Metric{
		reconciled: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "goclauded_reconciled_total",
				Help: "Documents reconcile pushed to goqueryd or failed to push, by collection and outcome",
			},
			[]string{constant.Collection, constant.Outcome},
		),
		indexed: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "goclauded_indexed_documents",
				Help: "Documents current in goqueryd after the last complete reconcile, by collection",
			},
			[]string{constant.Collection},
		),
		delivered: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name:    "goclauded_delivery_characters",
				Help:    "Characters of each rendered hook delivery context",
				Buckets: []float64{500, 1000, 1500, 2000, 2500, 3000, 4000},
			},
		),
		overBudget: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "goclauded_delivery_over_budget_total",
				Help: "Hook delivery contexts rendered over the delivery budget",
			},
		),
	}
	r.MustRegister(m.reconciled, m.indexed, m.delivered, m.overBudget)

	return m
}
