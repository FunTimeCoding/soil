package service_tester

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"

func WithAlertmanager(c *mock_client.Client) Option {
	return func(o *Tester) {
		o.Alertmanager = c
	}
}
