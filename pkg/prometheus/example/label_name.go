package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/prometheus"
)

func LabelName() {
	fmt.Println("Label names")

	for _, l := range prometheus.NewEnvironment().MustLabelNames(
		[]string{},
		constant.StartOfTime,
	).Values {
		fmt.Printf("  %s\n", l)
	}
}
