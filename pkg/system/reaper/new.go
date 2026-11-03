package reaper

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	reporter face.Reporter,
	y *prometheus.Registry,
) *Reaper {
	result := &Reaper{
		reporter: reporter,
		reaped:   NewCounter(),
		stop:     make(chan struct{}),
	}
	result.zombies = result.scan
	y.MustRegister(result.reaped)

	return result
}
