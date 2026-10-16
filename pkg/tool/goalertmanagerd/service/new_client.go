package service

import (
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

func newClient(i *instance.Instance) (*alertmanager.Client, error) {
	var a []alertmanager.Option

	if l := i.Alertmanager.Link; l.Host != "" {
		a = append(a, alertmanager.WithLink(l.Host, l.Port, l.Secure))
	}

	var p []prometheus.Option

	if l := i.Prometheus.Link; l.Host != "" {
		p = append(p, prometheus.WithLink(l.Host, l.Port, l.Secure))
	}

	return alertmanager.New(
		i.Alertmanager.Host,
		i.Alertmanager.Port,
		i.Alertmanager.Secure,
		i.Alertmanager.User,
		i.Alertmanager.Password,
		prometheus.New(
			i.Prometheus.Host,
			i.Prometheus.Port,
			i.Prometheus.Secure,
			i.Prometheus.User,
			i.Prometheus.Password,
			p...,
		),
		a...,
	), nil
}
