package peered

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
)

type Peered struct {
	alertmanager      face.AlertmanagerClient
	prometheus        face.PrometheusClient
	alertmanagerGuard *peer.Guard
	prometheusGuard   *peer.Guard
}
