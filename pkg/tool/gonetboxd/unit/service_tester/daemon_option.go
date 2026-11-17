package service_tester

import (
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/option"
	"github.com/prometheus/client_golang/prometheus"
)

func (o *Tester) daemonOption() *option.Netbox {
	result := option.New()
	result.LitePath = o.litePath
	result.PowerIntent = o.powerIntent
	result.Source = func(prometheus.ObserverVec) face.NetBoxClient {
		return o.client
	}

	return result
}
