package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func LabelName() {
	console.Line("Label names")

	for _, l := range prometheus.NewEnvironment().MustLabelNames(
		[]string{},
		constant.StartOfTime,
	).Values {
		console.Format("  %s\n", l)
	}
}
