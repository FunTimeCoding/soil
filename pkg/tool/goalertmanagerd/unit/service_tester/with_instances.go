package service_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
)

func WithInstances(names ...string) Option {
	return func(o *Tester) {
		o.instances = map[string]*Clients{}

		for _, n := range names {
			o.instances[n] = &Clients{
				Alertmanager: mock_client.New(o.t),
				Prometheus:   prometheusMock.New(o.t),
			}
		}
	}
}
