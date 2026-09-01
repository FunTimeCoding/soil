package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func Label() {
	console.Line("Labels:")

	for _, m := range prometheus.NewEnvironment().AllLabels() {
		console.Format("  %s\n", m)
	}
}
