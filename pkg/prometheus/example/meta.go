package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"slices"
)

var Groups = []string{
	"alertmanager",
	"events",
	"go",
	"kube",
	"net",
	"node",
	"pg",
	"postgres",
	"probe",
	"process",
	"prometheus",
	"promhttp",
	"pushgateway",
	"reloader",
}

func Meta() {
	c := prometheus.NewEnvironment()
	fmt.Println("Metadata")
	m := make(map[string][]string)

	for _, n := range c.AllMetrics() {
		for k, elements := range c.MustMetadata(n) {
			fmt.Printf("  %s\n", k)
			prefix, _ := key_value.Underscore(k)

			if slices.Contains(Groups, prefix) {
				m[prefix] = append(m[prefix], k)
			} else {
				m["other"] = append(m["other"], k)
			}

			for _, d := range elements {
				fmt.Printf("    %s", d.Type)

				if d.Unit != "" {
					fmt.Printf(" %s", d.Unit)
				}

				fmt.Printf(" %s\n", d.Help)
			}
		}
	}

	fmt.Println("Metric groups")

	for k, v := range m {
		fmt.Printf("%s: %d\n", k, len(v))

		for _, n := range v {
			fmt.Printf("  %s\n", n)
		}
	}
}
