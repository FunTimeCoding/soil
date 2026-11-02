package option

import (
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

func BuildPrometheus(i *instance.Instance) (face.Client, error) {
	var o []prometheus.Option

	if l := i.Prometheus.Link; l.Host != "" {
		o = append(o, prometheus.WithLink(l.Host, l.Port, l.Secure))
	}

	return prometheus.New(
		i.Prometheus.Host,
		i.Prometheus.Port,
		i.Prometheus.Secure,
		i.Prometheus.User,
		i.Prometheus.Password,
		o...,
	), nil
}
