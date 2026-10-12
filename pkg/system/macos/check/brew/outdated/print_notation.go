package outdated

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"github.com/funtimecoding/soil/pkg/system/macos/brew/formula"
	"github.com/funtimecoding/soil/pkg/system/macos/check/brew/outdated/option"
)

func printNotation(
	v []*formula.Formula,
	o *option.Outdated,
) {
	r := report.New()

	for _, e := range report.Trim(v, r, o.All, monitorConstant.GoBrew) {
		r.AddItem(
			monitorConstant.GoBrew,
			e.MonitorIdentifier,
			soil.Warning,
			e.CurrentVersion,
			e.Link,
			nil,
		)
	}

	r.Print()
}
