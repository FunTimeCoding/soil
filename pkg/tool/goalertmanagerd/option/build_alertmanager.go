package option

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
)

func BuildAlertmanager(i *instance.Instance) (face.Client, error) {
	var o []alertmanager.Option

	if l := i.Alertmanager.Link; l.Host != "" {
		o = append(o, alertmanager.WithLink(l.Host, l.Port, l.Secure))
	}

	return alertmanager.New(
		i.Alertmanager.Host,
		i.Alertmanager.Port,
		i.Alertmanager.Secure,
		i.Alertmanager.User,
		i.Alertmanager.Password,
		o...,
	), nil
}
