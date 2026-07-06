package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
)

func Statistic() {
	c := loki.NewEnvironment(true)
	// TODO: Somehow just returns zeroes
	fmt.Printf("Statistic: %s", c.Statistic(`{namespace!=""}`))
}
