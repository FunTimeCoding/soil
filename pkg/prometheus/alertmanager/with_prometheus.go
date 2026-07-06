package alertmanager

import "github.com/funtimecoding/soil/pkg/prometheus"

func WithPrometheus(p *prometheus.Client) Option {
	return func(c *Client) {
		c.prometheus = p
	}
}
