package alert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/check/alert/option"
	soil "github.com/funtimecoding/soil/pkg/constant"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
)

func printNotation(
	v []*alert.Alert,
	o *option.Alert,
) {
	r := report.New()

	for _, e := range report.Trim(v, r, o.All, monitorConstant.GoGenie) {
		var s soil.Severity

		if e.Acknowledged {
			s = soil.Warning
		} else {
			s = soil.Critical
		}

		r.AddItem(
			monitorConstant.GoGenie,
			e.MonitorIdentifier,
			s,
			e.Name,
			e.Link,
			&e.Create,
		)
	}

	r.Print()
}
