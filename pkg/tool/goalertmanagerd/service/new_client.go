package service

import (
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

func newClient(i *instance.Instance) (*alertmanager.Client, error) {
	return alertmanager.New(
		i.Alertmanager.Host,
		i.Alertmanager.Secure,
		i.Alertmanager.User,
		i.Alertmanager.Password,
		prometheus.New(
			i.Prometheus.Host,
			i.Prometheus.Port,
			i.Prometheus.Secure,
			i.Prometheus.User,
			i.Prometheus.Password,
			"",
		),
	), nil
}
