package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"slices"
)

func Meta() {
	c := prometheus.NewEnvironment()
	console.Line("Metadata")
	m := make(map[string][]string)

	for _, n := range c.AllMetrics() {
		for k, elements := range c.MustMetadata(n) {
			console.Format("  %s\n", k)
			prefix, _ := key_value.Underscore(k)

			if slices.Contains(constant.ExampleGroups, prefix) {
				m[prefix] = append(m[prefix], k)
			} else {
				m["other"] = append(m["other"], k)
			}

			for _, d := range elements {
				console.Format("    %s", d.Type)

				if d.Unit != "" {
					console.Format(" %s", d.Unit)
				}

				console.Format(" %s\n", d.Help)
			}
		}
	}

	console.Line("Metric groups")

	for k, v := range m {
		console.Format("%s: %d\n", k, len(v))

		for _, n := range v {
			console.Format("  %s\n", n)
		}
	}
}
