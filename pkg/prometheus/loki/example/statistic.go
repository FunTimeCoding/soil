package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
)

func Statistic() {
	c := loki.NewEnvironment(true)
	defer c.Close()
	console.Format("Statistic: %s", c.MustStatistic(`{namespace!=""}`))
}
