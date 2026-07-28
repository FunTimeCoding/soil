package outdated

import (
	"github.com/funtimecoding/soil/pkg/constant"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"github.com/funtimecoding/soil/pkg/system/macos/brew/formula"
	"github.com/funtimecoding/soil/pkg/system/macos/check/brew/outdated/option"
)

func printNotation(
	v []*formula.Formula,
	o *option.Outdated,
) {
	r := report.New()

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		monitor.GoBrew,
	) {
		r.AddItem(
			monitor.GoBrew,
			e.MonitorIdentifier,
			constant.Warning,
			e.CurrentVersion,
			e.Link,
			nil,
		)
	}

	r.Print()
}
