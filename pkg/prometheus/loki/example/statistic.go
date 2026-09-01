package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
)

func Statistic() {
	c := loki.NewEnvironment(true)
	// TODO: Somehow just returns zeroes
	console.Format("Statistic: %s", c.Statistic(`{namespace!=""}`))
}
