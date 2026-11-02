package peered

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
)

func New(
	a face.AlertmanagerClient,
	p face.PrometheusClient,
	alertmanagerGuard *peer.Guard,
	prometheusGuard *peer.Guard,
) *Peered {
	return &Peered{
		alertmanager:      a,
		prometheus:        p,
		alertmanagerGuard: alertmanagerGuard,
		prometheusGuard:   prometheusGuard,
	}
}
