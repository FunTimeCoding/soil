package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func Metric() {
	fmt.Println("Metric")

	for _, m := range prometheus.NewEnvironment().AllMetrics() {
		fmt.Printf("  %s\n", m)
	}
}
