package outdated

import (
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	item "github.com/funtimecoding/soil/pkg/monitor/item/constant"
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
		item.GoBrew,
	) {
		r.AddItem(
			item.GoBrew,
			e.MonitorIdentifier,
			constant.Warning,
			e.CurrentVersion,
			e.Link,
			nil,
		)
	}

	r.Print()
}
