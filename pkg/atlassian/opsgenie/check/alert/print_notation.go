package alert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/check/alert/option"
	"github.com/funtimecoding/soil/pkg/constant"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
)

func printNotation(
	v []*alert.Alert,
	o *option.Alert,
) {
	r := report.New()

	for _, e := range report.Trim(v, r, o.All, monitor.GoGenie) {
		var s constant.Severity

		if e.Acknowledged {
			s = constant.Warning
		} else {
			s = constant.Critical
		}

		r.AddItem(
			monitor.GoGenie,
			e.MonitorIdentifier,
			s,
			e.Name,
			e.Link,
			&e.Create,
		)
	}

	r.Print()
}
