package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func Label() {
	console.Line("Labels:")
	c := prometheus.NewEnvironment()
	defer c.Close()

	for _, m := range c.MustAllLabels() {
		console.Format("  %s\n", m)
	}
}
