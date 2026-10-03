package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
)

func Statistic() {
	c := loki.NewEnvironment(true)
	console.Format("Statistic: %s", c.Statistic(`{namespace!=""}`))
}
