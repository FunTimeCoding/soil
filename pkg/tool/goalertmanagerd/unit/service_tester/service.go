package service_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	o.service = pooled(
		o.t,
		o.Inventory(),
		func(i *instance.Instance) face.AlertmanagerClient {
			return o.Instance(i.Name).Alertmanager
		},
		func(i *instance.Instance) face.PrometheusClient {
			return o.Instance(i.Name).Prometheus
		},
		o.Reporter,
	)

	return o.service
}
