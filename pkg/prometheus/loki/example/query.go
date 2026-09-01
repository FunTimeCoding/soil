package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
)

func Query() {
	c := loki.NewEnvironment(false)
	console.Format("Query: %+v\n", c.Query(`rate({namespace="example"}[5m])`))
}
