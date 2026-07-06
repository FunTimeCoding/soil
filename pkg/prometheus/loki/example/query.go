package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus/loki"
)

func Query() {
	c := loki.NewEnvironment(false)
	fmt.Printf(
		"Query: %+v\n",
		c.Query(`rate({namespace="i9n"}[5m])`),
	)
}
