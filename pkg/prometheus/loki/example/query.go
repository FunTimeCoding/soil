package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
)

func Query() {
	c := loki.NewEnvironment(false)
	defer c.Close()
	console.Format(
		"Query: %+v\n",
		c.MustQuery(`rate({namespace="example"}[5m])`),
	)
}
