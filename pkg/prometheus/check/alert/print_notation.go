package alert

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	monitorConstant "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/check/alert/option"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func printNotation(
	v []*alert.Alert,
	o *option.Alert,
) {
	r := report.New()
	var relevant []*alert.Alert

	for _, e := range v {
		if !o.All && e.Severity == constant.InformationSeverity {
			continue
		}

		relevant = append(relevant, e)
	}

	for _, e := range report.Trim(relevant, r, o.All, monitorConstant.GoAlert) {
		var s soil.Severity

		if e.Severity == constant.CriticalSeverity {
			s = soil.Critical
		} else {
			s = soil.Warning
		}

		r.AddItem(
			monitorConstant.GoAlert,
			e.MonitorIdentifier,
			s,
			e.Name,
			e.Link,
			e.Start,
		)
	}

	r.Print()
}
