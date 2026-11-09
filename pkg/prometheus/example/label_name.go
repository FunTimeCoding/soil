package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"time"
)

func LabelName() {
	console.Line("Label names")
	c := prometheus.NewEnvironment()
	defer c.Close()

	for _, l := range c.MustLabelNames(
		[]string{},
		soil.Epoch,
		time.Time{},
	).Values {
		console.Format("  %s\n", l)
	}
}
