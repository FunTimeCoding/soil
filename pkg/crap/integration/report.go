package integration

import (
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/option"
)

func report(root string) *option.Report {
	o := option.New()
	o.Root = root
	o.Patterns = []string{constant.AllPackages}
	o.Threshold = constant.DefaultThreshold
	o.Tolerance = constant.DefaultTolerance
	o.Notation = true

	return o
}
