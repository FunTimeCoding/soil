package option

import (
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"github.com/prometheus/client_golang/prometheus"
)

type Netbox struct {
	option.Web
	Metric          option.Web
	LitePath        string
	PostgresLocator string
	PowerIntent     bool
	Source          func(prometheus.ObserverVec) face.NetBoxClient
}
