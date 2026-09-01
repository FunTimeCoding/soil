package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func Metric() {
	console.Line("Metric")

	for _, m := range prometheus.NewEnvironment().AllMetrics() {
		console.Format("  %s\n", m)
	}
}
