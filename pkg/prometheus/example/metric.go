package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func Metric() {
	console.Line("Metric")
	c := prometheus.NewEnvironment()
	defer c.Close()

	for _, m := range c.MustAllMetrics() {
		console.Format("  %s\n", m)
	}
}
