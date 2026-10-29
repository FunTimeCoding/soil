package option

import (
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/prometheus/client_golang/prometheus"
)

func New() *Netbox {
	return &Netbox{
		Source: func(o prometheus.ObserverVec) face.Client {
			return netbox.NewEnvironment(netbox.WithRequestDuration(o))
		},
	}
}
