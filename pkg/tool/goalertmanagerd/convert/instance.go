package convert

import "github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"

func Instance(
	i *instance.Instance,
	active bool,
) *SlimInstance {
	return &SlimInstance{
		Name:             i.Name,
		AlertmanagerHost: i.Alertmanager.Host,
		PrometheusHost:   i.Prometheus.Host,
		Active:           active,
	}
}
