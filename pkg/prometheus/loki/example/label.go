package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
	"time"
)

func Label() {
	c := loki.NewEnvironment(false)
	end := time.Now()
	start := end.AddDate(0, 0, -7)

	for _, l := range c.Labels(start, end) {
		fmt.Printf("Label: %s\n", l)
		fmt.Printf("  Values: %+v\n", c.LabelValues(start, end, l))
	}
}
