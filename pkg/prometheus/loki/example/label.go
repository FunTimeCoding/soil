package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	"time"
)

func Label() {
	c := loki.NewEnvironment(false)
	end := time.Now()
	start := end.AddDate(0, 0, -7)

	for _, l := range c.Labels(start, end) {
		console.Format("Label: %s\n", l)
		console.Format("  Values: %+v\n", c.LabelValues(start, end, l))
	}
}
