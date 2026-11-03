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
	}
	r.MustRegister(m.reconciled, m.indexed)

	return m
}
