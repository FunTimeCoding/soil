package netbox

import "github.com/prometheus/client_golang/prometheus"

func WithRequestDuration(o prometheus.ObserverVec) Option {
	return func(c *Client) {
		c.observer = o
	}
}
