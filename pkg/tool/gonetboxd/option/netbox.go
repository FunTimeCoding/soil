package option

import (
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/prometheus/client_golang/prometheus"
)

type Netbox struct {
	Address         string
	MetricAddress   string
	ServiceTokens   []string
	LitePath        string
	PostgresLocator string
	PowerIntent     bool
	Source          func(prometheus.ObserverVec) face.Client
}
