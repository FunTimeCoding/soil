package peered

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	prometheusFace "github.com/funtimecoding/soil/pkg/prometheus/face"
)

type Peered struct {
	alertmanager      face.Client
	prometheus        prometheusFace.Client
	alertmanagerGuard *peer.Guard
	prometheusGuard   *peer.Guard
}
