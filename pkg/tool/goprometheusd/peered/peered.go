package peered

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
)

type Peered struct {
	client face.PrometheusClient
	guard  *peer.Guard
}
