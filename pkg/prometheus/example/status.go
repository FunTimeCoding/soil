package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func Status() {
	c := prometheus.NewEnvironment()
	defer c.Close()
	console.Format("Status: %+v\n", c.MustStatus())
}
