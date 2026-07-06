package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func Status() {
	c := prometheus.NewEnvironment()
	fmt.Printf("Status: %+v\n", c.MustStatus())
}
