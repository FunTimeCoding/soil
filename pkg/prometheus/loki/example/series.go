package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
)

func Series() {
	c := loki.NewEnvironment(false)
	console.Format("Series: %s\n", c.Series(`{namespace="bot"}`))
}
