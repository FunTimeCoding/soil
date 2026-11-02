package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func LabelName() {
	console.Line("Label names")
	c := prometheus.NewEnvironment()
	defer c.Close()

	for _, l := range c.MustLabelNames(
		[]string{},
		soil.StartOfTime,
	).Values {
		console.Format("  %s\n", l)
	}
}
