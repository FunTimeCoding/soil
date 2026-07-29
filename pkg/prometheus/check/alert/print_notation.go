package alert

import (
	"github.com/funtimecoding/soil/pkg/constant"
	monitor "github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/monitor/report"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/check/alert/option"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func printNotation(
	v []*alert.Alert,
	o *option.Alert,
) {
	r := report.New()
	var relevant []*alert.Alert

	for _, e := range v {
		if !o.All && e.Severity == prometheus.InformationSeverity {
			continue
		}

		relevant = append(relevant, e)
	}

	for _, e := range report.Trim(
		relevant,
		r,
		o.All,
		monitor.GoAlert,
	) {
		var s constant.Severity

		if e.Severity == prometheus.CriticalSeverity {
			s = constant.Critical
		} else {
			s = constant.Warning
		}

		r.AddItem(
			monitor.GoAlert,
			e.MonitorIdentifier,
			s,
			e.Name,
			e.Link,
			e.Start,
		)
	}

	r.Print()
}
